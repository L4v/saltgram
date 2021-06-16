package handlers

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
	saltdata "saltgram/data"
	"saltgram/protos/auth/prauth"
	"saltgram/protos/email/premail"
	"saltgram/protos/users/prusers"
	"time"

	"github.com/gorilla/mux"
)

func (a *Auth) CheckPermissions(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.l.Printf("[ERROR] reading body: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	route := string(body)
	jws, _ := getUserJWS(r)
	_, err = a.ac.CheckPermissions(context.Background(),
		&prauth.PermissionRequest{
			Jws:    jws,
			Path:   route,
			Method: r.Method,
		})
	if err != nil {
		a.l.Printf("[ERROR] authenticating: %v\n", err)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Write([]byte("200 - OK"))
}

func (e *Email) ConfirmReset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	res, err := e.ec.ConfirmReset(context.Background(), &premail.ConfirmRequest{Token: token})
	if err != nil {
		e.l.Printf("[ERROR] confirming reset: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "emailforreset",
		Value:    res.Email,
		Expires:  time.Now().UTC().AddDate(0, 6, 0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/users",
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("200 - OK"))
}

func (e *Email) Activate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	_, err := e.ec.Activate(context.Background(), &premail.ActivateRequest{Token: token})
	if err != nil {
		e.l.Printf("[ERROR] activating email: %v\n", err)
		http.Error(w, "Failed to activate email", http.StatusBadRequest)
		return
	}

	w.Write([]byte("200 - OK"))
}

func (u *Users) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	jws, err := getUserJWS(r)
	if err != nil {
		u.l.Println("[ERROR] JWS not found")
		http.Error(w, "JWS not found", http.StatusBadRequest)
		return
	}

	token, err := jwt.ParseWithClaims(
		jws,
		&saltdata.AccessClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		},
	)

	if err != nil {
		u.l.Printf("[ERROR] parsing claims: %v", err)
		http.Error(w, "Error parsing claims", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*saltdata.AccessClaims)

	if !ok {
		u.l.Println("[ERROR] unable to parse claims")
		http.Error(w, "Error parsing claims: ", http.StatusInternalServerError)
		return
	}

	username := claims.Username

	dto := saltdata.ProflieDTO{}
	err = saltdata.FromJSON(&dto, r.Body)
	if err != nil {
		u.l.Printf("[ERROR] deserializing user data: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = dto.Validate()
	if err != nil {
		u.l.Printf("[ERROR] validating user data: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err = u.uc.UpdateProfile(context.Background(), &prusers.UpdateRequest{
		OldUsername: username,
		NewUsername: dto.Username,
		Email:       dto.Email,
		FullName:    dto.FullName,
		Public:      !dto.PrivateProfile,
		Taggable:    dto.Taggable,
		Description: dto.Description,
		PhoneNumber: dto.PhoneNumber,
		Gender: dto.Gender,
		DateOfBirth: dto.DateOfBirth.Unix(),
		WebSite: dto.WebSite,
		PrivateProfile: dto.PrivateProfile,
	})


	if err != nil {
		u.l.Printf("[ERROR] updating profile: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Updated %s"))

}

func (a *Auth) UpdateJWTUsername(w http.ResponseWriter, r *http.Request) {

	// TODO: PART 1 JWS
	jws, err := getUserJWS(r)
	if err != nil {
		a.l.Println("[ERROR] JWS not found")
		http.Error(w, "JWS not found", http.StatusBadRequest)
		return
	}

	token, err := jwt.ParseWithClaims(
		jws,
		&saltdata.AccessClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		},
	)

	if err != nil {
		a.l.Printf("[ERROR] parsing claims: %v", err)
		http.Error(w, "Error parsing claims", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*saltdata.AccessClaims)

	if !ok {
		a.l.Println("[ERROR] unable to parse claims")
		http.Error(w, "Error parsing claims: ", http.StatusInternalServerError)
		return
	}

	username := claims.Username
	password := claims.Password

	type NewUsernameDto struct {
		NewUsername string `json:"newUsername" validate:"required"`
	}

	// TODO: PART 2 UPDATING

	dto := NewUsernameDto{}
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		a.l.Printf("[ERROR] deserializing JWTDto: %v", err)
		http.Error(w, "Failed to deserialize JWTDto", http.StatusBadRequest)
		return
	}

	_, err = a.ac.UpdateRefresh(context.Background(), &prauth.UpdateRefreshRequest{OldUsername: username, NewUsername: dto.NewUsername})
	if err != nil {
		a.l.Printf("[ERROR] updating jwt: %v\n", err)
		http.Error(w, "Failed to update jwt", http.StatusBadRequest)
		return
	}

	// TODO: PART 3 SETTING NEW UPDATED TOKEN

	res, err := a.ac.GetJWT(context.Background(), &prauth.JWTRequest{Username: dto.NewUsername, Password: password})
	if err != nil {
		a.l.Printf("[ERROR] getting new updated jwt: %v\n", err)
		http.Error(w, "Failed to get new updated jwt", http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "refresh",
		Value:    res.Refresh,
		Expires:  time.Now().UTC().AddDate(0, 6, 0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(res.Jws))

}

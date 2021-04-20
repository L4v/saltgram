package handlers

import (
	"net/http"
	"saltgram/data"
	"time"

	"github.com/gorilla/mux"
)

func (e *Emails) ConfirmReset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	email, err := data.ConfirmPasswordReset(token)
	if err != nil {
		e.l.Printf("[ERROR] confirming password reset: %v\n", err)
		http.Error(w, "Failed to confirm password reset", http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     "email",
		Value:    email,
		Expires:  time.Now().UTC().AddDate(0, 6, 0),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Activated"))
}

func (e *Emails) Activate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	token := vars["token"]

	err := data.ActivateEmail(token)
	if err != nil {
		e.l.Printf("[ERROR] activating email: %v", err)
		http.Error(w, "Failed activating email: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Email activated"))
}

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getUserID(r)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	u.l.Println("Handling PUT Users", id)
	// NOTE(Jovan): Safe to cast because middleware makes sure nothing's wrong
	user := r.Context().Value(KeyUser{}).(data.User)

	err = data.UpdateUser(id, &user)
	if err == data.ErrUserNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not updated: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
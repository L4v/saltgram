package data

import "time"

type RequestStatus string

const (
	PENDING RequestStatus = "PENDING"
	ACCEPTED
	DENIED
)

type Profile struct {
	UserID         uint64 `gorm:"primaryKey"`
	Username       string `json:"username" gorm:"unique"`
	User           User
	Public         bool       `json:"isPublic"`
	Taggable       bool       `json:"isTaggable"`
	Description    string     `json:"description"`
	Following      []*Profile `gorm:"many2many:profile_following;"`
	Profiles       []FollowRequest
	Requests       []FollowRequest
	PhoneNumber    string    `json:"phoneNumber"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	WebSite        string    `json:"webSite"`
	PrivateProfile bool      `json:"privateProfile"`
}

type FollowRequest struct {
	ID            uint64        `json:"-"`
	ProfileID     uint64        `json:"profileId"`
	RequestID     uint64        `json:"followerId"`
	RequestStatus RequestStatus `json:"stats"`
}

//TODO(Marko add profliPicture?)
type ProfileFollowerDTO struct {
	Username string
	FullName string
}

func (db *DBConn) GetProfiles() []*Profile {
	profiles := []*Profile{}
	db.DB.Find(&profiles)
	return profiles
}

func (db *DBConn) AddProfile(p *Profile) error {
	// err := p.User.GenerateSaltAndHashedPassword()
	// if err != nil {
	// 	return err
	// }
	return db.DB.Create(p).Error
}

func (db *DBConn) UpdateProfile(p *Profile) error {
	return db.DB.Save(&p).Error
}

func (db *DBConn) GetProfileByUsername(username string) (*Profile, error) {
	profile := Profile{}
	err := db.DB.Where("username = ?", username).First(&profile).Error
	return &profile, err
}

func CheckIfFollowing(db *DBConn, profile_username string, following_user_id uint64) (bool, error) {
	var count int64
	// err := db.DB.Table("profile_followers").Where("profile_username = ? AND follower_username = ?", profile, username).Count(&count).Error
	err := db.DB.Raw("SELECT COUNT(*) FROM profile_following LEFT JOIN profiles on profile_user_id = user_id WHERE username = ? AND following_user_id = ? ", profile_username, following_user_id).Count(&count).Error
	if err != nil {
		return false, err
	}
	exists := count > 0
	return exists, nil
}

func GetFollowerCount(db *DBConn, username string) (int64, error) {
	var count int64
	// err := db.DB.Table("profile_followers").Where("follower_username = ?", username).Count(&count).Error
	err := db.DB.Raw("SELECT COUNT(*) FROM profile_following LEFT JOIN profiles on following_user_id = user_id WHERE username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	// if count == 0 {
	// 	return -1, nil
	// }
	return count, nil
}

func GetFollowingCount(db *DBConn, username string) (int64, error) {
	var count int64
	// err := db.DB.Table("profile_followers").Where("profile_username = ?", username).Count(&count).Error
	err := db.DB.Raw("SELECT COUNT(*) FROM profile_following LEFT JOIN profiles on profile_user_id = user_id WHERE username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	// if count == 0 {
	// 	return -1, nil
	// }
	return count, nil
}

func SetFollow(db *DBConn, profile *Profile, profileToFollow *Profile) error {
	//db.DB.Model(&profile).Association("Following").Append(&profileToFollow)
	profile.Following = append(profile.Following, profileToFollow)
	return db.DB.Save(&profile).Error
}

func CreateFollowRequest(db *DBConn, profile *Profile, request *Profile) error {
	//profile.Requests = append(profile.Requests, request)
	fr := FollowRequest{
		ProfileID:     profile.UserID,
		RequestID:     request.UserID,
		RequestStatus: PENDING,
	}
	return db.DB.Create(&fr).Error
}

func GetFollowers(db *DBConn, profile *Profile) ([]Profile, error) {
	// TODO This does not work. Need to move to graph db
	//
	var followers []Profile
	err := db.DB.Where("following_user_id = ?", profile.UserID).Association("Following").Find(&followers)
	//err := db.DB.Preload("Followers").Where("follower_username = ?", username).Find(&followers).Error
	//err := db.DB.Raw("SELECT * FROM profile_following LEFT JOIN profiles on following_user_id = user_id WHERE username = ?", username).Scan(&followers).Error
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func GetFollowing(db *DBConn, profile *Profile) ([]Profile, error) {
	// TODO This does not work. Need to move to graph db
	//
	var following []Profile
	//err := db.DB.Raw("SELECT * FROM profile_following LEFT JOIN profiles on profile_user_id = user_id WHERE username = ?", username).Scan(&following).Error
	err := db.DB.Model(&profile).Where("profile_user_id = ?", profile.UserID).Association("Following").Find(&following)
	if err != nil {
		return nil, err
	}
	return following, nil
}

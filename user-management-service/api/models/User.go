package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string  `json:"id" gorm:"primaryKey"`
	Email        string  `json:"email" gorm:"unique"`
	Name         string  `json:"name"`
	Password     *string `json:"password"`
	GoogleID     *string `json:"google_id"`
	AppleID      *string `json:"apple_id"`
	MicrosoftID  *string `json:"microsoft_id"`
	GoogleAvatar *string `json:"google_avatar"`

	GoogleAccessToken  *string    `json:"google_access_token"`
	GoogleRefreshToken *string    `json:"google_refresh_token"`
	GoogleTokenExpiry  *time.Time `json:"google_token_expiry"`

	AppleAccessToken  *string    `json:"apple_access_token"`
	AppleRefreshToken *string    `json:"apple_refresh_token"`
	AppleTokenExpiry  *time.Time `json:"apple_token_expiry"`

	MicrosoftAccessToken  *string    `json:"microsoft_access_token"`
	MicrosoftRefreshToken *string    `json:"microsoft_refresh_token"`
	MicrosoftTokenExpiry  *time.Time `json:"microsoft_token_expiry"`

	Role      string    `json:"role" gorm:"default:'attendee'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

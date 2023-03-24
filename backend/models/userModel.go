package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string     `gorm:"type:varchar(100);not null"`
	Email    string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string     `gorm:"type:varchar(100);not null;" json:"-"`
	// Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	// Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	// Photo     *string    `gorm:"not null;default:'default.png'"`
	// Verified  *bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type SignUpInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	// @todo uncomment this
	// Password        string `json:"password" validate:"required,min=8"`
	Password string `json:"password" validate:"required"`
	// PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	// Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refreshToken"  validate:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
	// Role      string    `json:"role,omitempty"`
	// Photo     string    `json:"photo,omitempty"`
	// Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:    *user.ID,
		Name:  user.Name,
		Email: user.Email,
		// Role:      *user.Role,
		// Photo:     *user.Photo,
		// Provider:  *user.Provider,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}

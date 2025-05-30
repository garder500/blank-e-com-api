package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"uniqueIndex;not null" json:"username"`
	Password string   `gorm:"not null" json:"-"`
	ApiKeys  []ApiKey `gorm:"foreignKey:UserID" json:"api_keys"`
	Roles    []Role   `gorm:"many2many:user_roles;" json:"roles"`
}

type ApiKey struct {
	gorm.Model
	Key       string    `gorm:"uniqueIndex;not null" json:"key"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

type Role struct {
	gorm.Model
	Name        string        `gorm:"uniqueIndex;not null" json:"name"`
	Description string        `gorm:"not null" json:"description"`
	Users       []User        `gorm:"many2many:user_roles;" json:"users"`
	Permissions []Permissions `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permissions struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;" json:"roles"`
}

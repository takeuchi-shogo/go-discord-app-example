package entity

import "time"

const (
	UserModelName = "User"
	UserTableName = "users"
)

type User struct {
	ID          int64       `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserDetail  UserDetail  `gorm:"foreignKey:UserID;references:ID"`
	UserAccount UserAccount `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt   time.Time   `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;type:datetime;"`
}

// func (u User) ToModel() (*user.User, error) {}

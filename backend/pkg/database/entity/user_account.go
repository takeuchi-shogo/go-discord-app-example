package entity

import "time"

type UserAccount struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserID    int64     `gorm:"column:user_id;type:bigint;"`
	Email     string    `gorm:"column:email;type:varchar;size:255;uniqueIndex"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}

package entity

import "time"

const (
	UserDetailModelName = "UserDetail"
	UserDetailTableName = "user_detail"
)

type UserDetail struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserID    int64     `gorm:"column:user_id;type:bigint;"`
	Name      string    `gorm:"column:name;type:varchar;size:255;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}

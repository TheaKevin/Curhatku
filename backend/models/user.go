package models

import "time"

type UserTab struct {
	UserID    string    `json:"-" gorm:"not null; type: varchar(100); unique"`
	Username  string    `json:"username" gorm:"not null; type: varchar(100); unique"`
	Email     string    `json:"email" gorm:"not null; type: varchar(100); unique"`
	Password  []byte    `json:"-" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	BirthDate time.Time `json:"birth_date" gorm:"type: timestamp"`
}

func (m *UserTab) TableName() string {
	return "user_Tab"
}

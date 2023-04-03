package models

type CurhatTab struct {
	UserID        string `json:"user_id" gorm:"not null; type: varchar(100)"`
	CurhatID      string `json:"curhat_id" gorm:"not null; type: varchar(100); unique"`
	CurhatContent string `json:"curhat_content" gorm:"not null; type: varchar(255)"`
}

func (m *CurhatTab) TableName() string {
	return "curhat_tab"
}

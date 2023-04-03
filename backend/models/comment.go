package models

type CommentTab struct {
	UserID    string `json:"user_id" gorm:"not null; type: varchar(100)"`
	CurhatID  string `json:"curhat_id" gorm:"not null; type: varchar(100)"`
	CommentID string `json:"comment_id" gorm:"not null; type: varchar(100); unique"`
	Comment   string `json:"comment" gorm:"not null; type: varchar(255)"`
}

func (m *CommentTab) TableName() string {
	return "comment_tab"
}

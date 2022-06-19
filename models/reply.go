package models

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity        string         `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 唯一标识
	CommentIdentity string         `gorm:"column:comment_identity;type:varchar(36);" json:"comment_identity"` // 评论的唯一标识
	UserIdentity    string         `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       // 用户的唯一标识
	Reply           string         `gorm:"column:reply;type:varchar(255);" json:"reply"`                      //回复内容
}

func (table *Reply) TableName() string {
	return "reply"
}

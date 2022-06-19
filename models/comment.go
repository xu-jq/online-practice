package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity        string         `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 唯一标识
	ContentIdentity string         `gorm:"column:content_identity;type:varchar(36);" json:"content_identity"` // 内容（包括问题及题解）的唯一标识
	UserIdentity    string         `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       // 用户的唯一标识
	Comment         string         `gorm:"column:comment;type:varchar(255);" json:"comment"`                  //评论内容
}

func (table *Comment) TableName() string {
	return "comment"
}

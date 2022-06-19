package models

import (
	"gorm.io/gorm"
	"time"
)

type ProblemExplanation struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	Identity        string         `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 唯一标识
	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` // 问题的唯一标识
	UserIdentity    string         `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       // 用户的唯一标识
	Title           string         `gorm:"column:title;type:varchar(255);" json:"title"`                      //题解标题
	Content         string         `gorm:"column:content;type:text;" json:"content"`                          //题解内容
	ReadNum         int            `gorm:"column:read_num;type:int(11);" json:"read_num"`                     //阅读量
	LikeNum         int            `gorm:"column:like_num;type:int(11);" json:"like_num"`                     //点赞数
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
}

func (table *ProblemExplanation) TableName() string {
	return "problem_explanation"
}

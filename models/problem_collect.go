package models

import (
	"gorm.io/gorm"
	"time"
)

type ProblemCollect struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity        string         `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 唯一标识
	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` // 问题的唯一标识
	UserIdentity    string         `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       // 用户的唯一标识
}

func (table *ProblemCollect) TableName() string {
	return "problem_collect"
}

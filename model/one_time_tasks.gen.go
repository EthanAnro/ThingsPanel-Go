// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOneTimeTask = "one_time_tasks"

// OneTimeTask mapped from table <one_time_tasks>
type OneTimeTask struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	SceneAutomationID string    `gorm:"column:scene_automation_id;not null;comment:场景联动ID（外键-关联删除）" json:"scene_automation_id"`          // 场景联动ID（外键-关联删除）
	ExecutionTime     time.Time `gorm:"column:execution_time;not null;comment:执行时间" json:"execution_time"`                               // 执行时间
	ExecutingState    string    `gorm:"column:executing_state;not null;comment:1.执行状态 NEX-未执行 EXE-已执行 EXP-过期未执行" json:"executing_state"` // 1.执行状态 NEX-未执行 EXE-已执行 EXP-过期未执行
	Enabled           string    `gorm:"column:enabled;not null;comment:是否启用 Y-启用 N-停用" json:"enabled"`                                   // 是否启用 Y-启用 N-停用
	Remark            *string   `gorm:"column:remark" json:"remark"`
	ExpirationTime    int64     `gorm:"column:expiration_time;not null;comment:过期时间（默认大于执行时间五分钟5min10min30min1h1day）单位分钟" json:"expiration_time"` // 过期时间（默认大于执行时间五分钟5min10min30min1h1day）单位分钟
}

// TableName OneTimeTask's table name
func (*OneTimeTask) TableName() string {
	return TableNameOneTimeTask
}

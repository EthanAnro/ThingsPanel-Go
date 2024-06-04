// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAlarmHistory = "alarm_history"

// AlarmHistory mapped from table <alarm_history>
type AlarmHistory struct {
	ID                string    `gorm:"column:id;primaryKey" json:"id"`
	AlarmConfigID     string    `gorm:"column:alarm_config_id;not null" json:"alarm_config_id"`
	GroupID           string    `gorm:"column:group_id;not null" json:"group_id"`
	SceneAutomationID string    `gorm:"column:scene_automation_id;not null" json:"scene_automation_id"`
	Name              string    `gorm:"column:name;not null;comment:告警名称" json:"name"`                            // 告警名称
	Description       *string   `gorm:"column:description;comment:告警描述" json:"description"`                       // 告警描述
	Content           *string   `gorm:"column:content;comment:内容（什么原因导致的告警）" json:"content"`                      // 内容（什么原因导致的告警）
	AlarmStatus       string    `gorm:"column:alarm_status;not null;comment:L 底 M中 H 高 N 正常" json:"alarm_status"` // L 底 M中 H 高 N 正常
	TenantID          string    `gorm:"column:tenant_id;not null;comment:租户" json:"tenant_id"`                    // 租户
	Remark            *string   `gorm:"column:remark" json:"remark"`
	CreateAt          time.Time `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`                   // 创建时间
	AlarmDeviceList   string    `gorm:"column:alarm_device_list;not null;comment:触发设备id" json:"alarm_device_list"` // 触发设备id
}

// TableName AlarmHistory's table name
func (*AlarmHistory) TableName() string {
	return TableNameAlarmHistory
}

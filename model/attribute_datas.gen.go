// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAttributeData = "attribute_datas"

// AttributeData mapped from table <attribute_datas>
type AttributeData struct {
	ID       string    `gorm:"column:id;not null" json:"id"`
	DeviceID string    `gorm:"column:device_id;not null;comment:设备id（外键-关联删除）" json:"device_id"` // 设备id（外键-关联删除）
	Key      string    `gorm:"column:key;not null;comment:数据标识符" json:"key"`                     // 数据标识符
	T        time.Time `gorm:"column:ts;not null;comment:上报时间" json:"ts"`                        // 上报时间
	BoolV    *bool     `gorm:"column:bool_v" json:"bool_v"`
	NumberV  *float64  `gorm:"column:number_v" json:"number_v"`
	StringV  *string   `gorm:"column:string_v" json:"string_v"`
	TenantID *string   `gorm:"column:tenant_id" json:"tenant_id"`
}

// TableName AttributeData's table name
func (*AttributeData) TableName() string {
	return TableNameAttributeData
}

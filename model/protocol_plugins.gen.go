// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProtocolPlugin = "protocol_plugins"

// ProtocolPlugin mapped from table <protocol_plugins>
type ProtocolPlugin struct {
	ID             string    `gorm:"column:id;primaryKey;comment:Id" json:"id"`                                                      // Id
	Name           string    `gorm:"column:name;not null;comment:插件名称" json:"name"`                                                  // 插件名称
	DeviceType     int16     `gorm:"column:device_type;not null;default:1;comment:接入设备类型 (1-直连设备 2-网关设备 默认直连设备)" json:"device_type"` // 接入设备类型 (1-直连设备 2-网关设备 默认直连设备)
	ProtocolType   string    `gorm:"column:protocol_type;not null;comment:协议类型" json:"protocol_type"`                                // 协议类型
	AccessAddress  *string   `gorm:"column:access_address;comment:接入地址" json:"access_address"`                                       // 接入地址
	HTTPAddress    *string   `gorm:"column:http_address;comment:HTTP服务地址" json:"http_address"`                                       // HTTP服务地址
	SubTopicPrefix *string   `gorm:"column:sub_topic_prefix;comment:插件订阅前缀" json:"sub_topic_prefix"`                                 // 插件订阅前缀
	Description    *string   `gorm:"column:description;comment:描述" json:"description"`                                               // 描述
	AdditionalInfo *string   `gorm:"column:additional_info;comment:附加信息" json:"additional_info"`                                     // 附加信息
	CreatedAt      time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                      // 创建时间
	UpdateAt       time.Time `gorm:"column:update_at;not null;comment:更新时间" json:"update_at"`                                        // 更新时间
	Remark         *string   `gorm:"column:remark;comment:备注" json:"remark"`                                                         // 备注
}

// TableName ProtocolPlugin's table name
func (*ProtocolPlugin) TableName() string {
	return TableNameProtocolPlugin
}

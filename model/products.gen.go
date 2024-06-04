// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID             string    `gorm:"column:id;primaryKey;comment:uuid" json:"id"`                // uuid
	Name           string    `gorm:"column:name;not null;comment:产品名称" json:"name"`              // 产品名称
	Description    *string   `gorm:"column:description;comment:描述" json:"description"`           // 描述
	ProductType    *string   `gorm:"column:product_type;comment:产品类型" json:"product_type"`       // 产品类型
	ProductKey     *string   `gorm:"column:product_key;comment:产品key" json:"product_key"`        // 产品key
	ProductModel   *string   `gorm:"column:product_model;comment:产品型号(编号)" json:"product_model"` // 产品型号(编号)
	ImageURL       *string   `gorm:"column:image_url;comment:图片" json:"image_url"`               // 图片
	CreatedAt      time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`  // 创建时间
	Remark         *string   `gorm:"column:remark" json:"remark"`
	AdditionalInfo *string   `gorm:"column:additional_info" json:"additional_info"`
	TenantID       *string   `gorm:"column:tenant_id;comment:租户id" json:"tenant_id"` // 租户id
	DeviceConfigID *string   `gorm:"column:device_config_id" json:"device_config_id"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}

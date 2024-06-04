// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"project/model"
)

func newDeviceConfig(db *gorm.DB, opts ...gen.DOOption) deviceConfig {
	_deviceConfig := deviceConfig{}

	_deviceConfig.deviceConfigDo.UseDB(db, opts...)
	_deviceConfig.deviceConfigDo.UseModel(&model.DeviceConfig{})

	tableName := _deviceConfig.deviceConfigDo.TableName()
	_deviceConfig.ALL = field.NewAsterisk(tableName)
	_deviceConfig.ID = field.NewString(tableName, "id")
	_deviceConfig.Name = field.NewString(tableName, "name")
	_deviceConfig.DeviceTemplateID = field.NewString(tableName, "device_template_id")
	_deviceConfig.DeviceType = field.NewString(tableName, "device_type")
	_deviceConfig.ProtocolType = field.NewString(tableName, "protocol_type")
	_deviceConfig.VoucherType = field.NewString(tableName, "voucher_type")
	_deviceConfig.ProtocolConfig = field.NewString(tableName, "protocol_config")
	_deviceConfig.DeviceConnType = field.NewString(tableName, "device_conn_type")
	_deviceConfig.AdditionalInfo = field.NewString(tableName, "additional_info")
	_deviceConfig.Description = field.NewString(tableName, "description")
	_deviceConfig.TenantID = field.NewString(tableName, "tenant_id")
	_deviceConfig.CreatedAt = field.NewTime(tableName, "created_at")
	_deviceConfig.UpdatedAt = field.NewTime(tableName, "updated_at")
	_deviceConfig.Remark = field.NewString(tableName, "remark")
	_deviceConfig.OtherConfig = field.NewString(tableName, "other_config")

	_deviceConfig.fillFieldMap()

	return _deviceConfig
}

type deviceConfig struct {
	deviceConfigDo

	ALL              field.Asterisk
	ID               field.String // Id
	Name             field.String // 名称
	DeviceTemplateID field.String // 设备模板id
	DeviceType       field.String // 设备类型 1直连设备 2网关设备 3网关子设备
	ProtocolType     field.String // 协议类型
	VoucherType      field.String // 凭证类型
	ProtocolConfig   field.String // 协议表单配置
	DeviceConnType   field.String // 设备连接方式（默认A）A-设备连接平台B-平台连接设备
	AdditionalInfo   field.String // 附加信息
	Description      field.String // 描述
	TenantID         field.String // 租户id
	CreatedAt        field.Time   // 创建时间
	UpdatedAt        field.Time   // 更新时间
	Remark           field.String // 备注
	OtherConfig      field.String // 其他配置

	fieldMap map[string]field.Expr
}

func (d deviceConfig) Table(newTableName string) *deviceConfig {
	d.deviceConfigDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceConfig) As(alias string) *deviceConfig {
	d.deviceConfigDo.DO = *(d.deviceConfigDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceConfig) updateTableName(table string) *deviceConfig {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Name = field.NewString(table, "name")
	d.DeviceTemplateID = field.NewString(table, "device_template_id")
	d.DeviceType = field.NewString(table, "device_type")
	d.ProtocolType = field.NewString(table, "protocol_type")
	d.VoucherType = field.NewString(table, "voucher_type")
	d.ProtocolConfig = field.NewString(table, "protocol_config")
	d.DeviceConnType = field.NewString(table, "device_conn_type")
	d.AdditionalInfo = field.NewString(table, "additional_info")
	d.Description = field.NewString(table, "description")
	d.TenantID = field.NewString(table, "tenant_id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.Remark = field.NewString(table, "remark")
	d.OtherConfig = field.NewString(table, "other_config")

	d.fillFieldMap()

	return d
}

func (d *deviceConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceConfig) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 15)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["device_template_id"] = d.DeviceTemplateID
	d.fieldMap["device_type"] = d.DeviceType
	d.fieldMap["protocol_type"] = d.ProtocolType
	d.fieldMap["voucher_type"] = d.VoucherType
	d.fieldMap["protocol_config"] = d.ProtocolConfig
	d.fieldMap["device_conn_type"] = d.DeviceConnType
	d.fieldMap["additional_info"] = d.AdditionalInfo
	d.fieldMap["description"] = d.Description
	d.fieldMap["tenant_id"] = d.TenantID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["remark"] = d.Remark
	d.fieldMap["other_config"] = d.OtherConfig
}

func (d deviceConfig) clone(db *gorm.DB) deviceConfig {
	d.deviceConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceConfig) replaceDB(db *gorm.DB) deviceConfig {
	d.deviceConfigDo.ReplaceDB(db)
	return d
}

type deviceConfigDo struct{ gen.DO }

type IDeviceConfigDo interface {
	gen.SubQuery
	Debug() IDeviceConfigDo
	WithContext(ctx context.Context) IDeviceConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceConfigDo
	WriteDB() IDeviceConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceConfigDo
	Not(conds ...gen.Condition) IDeviceConfigDo
	Or(conds ...gen.Condition) IDeviceConfigDo
	Select(conds ...field.Expr) IDeviceConfigDo
	Where(conds ...gen.Condition) IDeviceConfigDo
	Order(conds ...field.Expr) IDeviceConfigDo
	Distinct(cols ...field.Expr) IDeviceConfigDo
	Omit(cols ...field.Expr) IDeviceConfigDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	Group(cols ...field.Expr) IDeviceConfigDo
	Having(conds ...gen.Condition) IDeviceConfigDo
	Limit(limit int) IDeviceConfigDo
	Offset(offset int) IDeviceConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceConfigDo
	Unscoped() IDeviceConfigDo
	Create(values ...*model.DeviceConfig) error
	CreateInBatches(values []*model.DeviceConfig, batchSize int) error
	Save(values ...*model.DeviceConfig) error
	First() (*model.DeviceConfig, error)
	Take() (*model.DeviceConfig, error)
	Last() (*model.DeviceConfig, error)
	Find() ([]*model.DeviceConfig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceConfig, err error)
	FindInBatches(result *[]*model.DeviceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeviceConfig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceConfigDo
	Assign(attrs ...field.AssignExpr) IDeviceConfigDo
	Joins(fields ...field.RelationField) IDeviceConfigDo
	Preload(fields ...field.RelationField) IDeviceConfigDo
	FirstOrInit() (*model.DeviceConfig, error)
	FirstOrCreate() (*model.DeviceConfig, error)
	FindByPage(offset int, limit int) (result []*model.DeviceConfig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceConfigDo) Debug() IDeviceConfigDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceConfigDo) WithContext(ctx context.Context) IDeviceConfigDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceConfigDo) ReadDB() IDeviceConfigDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceConfigDo) WriteDB() IDeviceConfigDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceConfigDo) Session(config *gorm.Session) IDeviceConfigDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceConfigDo) Clauses(conds ...clause.Expression) IDeviceConfigDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceConfigDo) Returning(value interface{}, columns ...string) IDeviceConfigDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceConfigDo) Not(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceConfigDo) Or(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceConfigDo) Select(conds ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceConfigDo) Where(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceConfigDo) Order(conds ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceConfigDo) Distinct(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceConfigDo) Omit(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceConfigDo) Join(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceConfigDo) Group(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceConfigDo) Having(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceConfigDo) Limit(limit int) IDeviceConfigDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceConfigDo) Offset(offset int) IDeviceConfigDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceConfigDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceConfigDo) Unscoped() IDeviceConfigDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceConfigDo) Create(values ...*model.DeviceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceConfigDo) CreateInBatches(values []*model.DeviceConfig, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceConfigDo) Save(values ...*model.DeviceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceConfigDo) First() (*model.DeviceConfig, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Take() (*model.DeviceConfig, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Last() (*model.DeviceConfig, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Find() ([]*model.DeviceConfig, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeviceConfig), err
}

func (d deviceConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeviceConfig, err error) {
	buf := make([]*model.DeviceConfig, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceConfigDo) FindInBatches(result *[]*model.DeviceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceConfigDo) Attrs(attrs ...field.AssignExpr) IDeviceConfigDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceConfigDo) Assign(attrs ...field.AssignExpr) IDeviceConfigDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceConfigDo) Joins(fields ...field.RelationField) IDeviceConfigDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceConfigDo) Preload(fields ...field.RelationField) IDeviceConfigDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceConfigDo) FirstOrInit() (*model.DeviceConfig, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceConfig), nil
	}
}

func (d deviceConfigDo) FirstOrCreate() (*model.DeviceConfig, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeviceConfig), nil
	}
}

func (d deviceConfigDo) FindByPage(offset int, limit int) (result []*model.DeviceConfig, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceConfigDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceConfigDo) Delete(models ...*model.DeviceConfig) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceConfigDo) withDO(do gen.Dao) *deviceConfigDo {
	d.DO = *do.(*gen.DO)
	return d
}

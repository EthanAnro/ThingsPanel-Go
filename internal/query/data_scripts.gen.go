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

	"project/internal/model"
)

func newDataScript(db *gorm.DB, opts ...gen.DOOption) dataScript {
	_dataScript := dataScript{}

	_dataScript.dataScriptDo.UseDB(db, opts...)
	_dataScript.dataScriptDo.UseModel(&model.DataScript{})

	tableName := _dataScript.dataScriptDo.TableName()
	_dataScript.ALL = field.NewAsterisk(tableName)
	_dataScript.ID = field.NewString(tableName, "id")
	_dataScript.Name = field.NewString(tableName, "name")
	_dataScript.DeviceConfigID = field.NewString(tableName, "device_config_id")
	_dataScript.EnableFlag = field.NewString(tableName, "enable_flag")
	_dataScript.Content = field.NewString(tableName, "content")
	_dataScript.ScriptType = field.NewString(tableName, "script_type")
	_dataScript.LastAnalogInput = field.NewString(tableName, "last_analog_input")
	_dataScript.Description = field.NewString(tableName, "description")
	_dataScript.CreatedAt = field.NewTime(tableName, "created_at")
	_dataScript.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dataScript.Remark = field.NewString(tableName, "remark")

	_dataScript.fillFieldMap()

	return _dataScript
}

type dataScript struct {
	dataScriptDo

	ALL             field.Asterisk
	ID              field.String // Id
	Name            field.String // 名称
	DeviceConfigID  field.String // 设备配置id 关联删除
	EnableFlag      field.String // 启用标志Y-启用 N-停用 默认启用
	Content         field.String // 内容
	ScriptType      field.String // 脚本类型 A-遥测上报预处理B-遥测下发预处理C-属性上报预处理D-属性下发预处理
	LastAnalogInput field.String // 上次模拟输入
	Description     field.String // 描述
	CreatedAt       field.Time   // 创建时间
	UpdatedAt       field.Time   // 更新时间
	Remark          field.String // 备注

	fieldMap map[string]field.Expr
}

func (d dataScript) Table(newTableName string) *dataScript {
	d.dataScriptDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dataScript) As(alias string) *dataScript {
	d.dataScriptDo.DO = *(d.dataScriptDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dataScript) updateTableName(table string) *dataScript {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "id")
	d.Name = field.NewString(table, "name")
	d.DeviceConfigID = field.NewString(table, "device_config_id")
	d.EnableFlag = field.NewString(table, "enable_flag")
	d.Content = field.NewString(table, "content")
	d.ScriptType = field.NewString(table, "script_type")
	d.LastAnalogInput = field.NewString(table, "last_analog_input")
	d.Description = field.NewString(table, "description")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.Remark = field.NewString(table, "remark")

	d.fillFieldMap()

	return d
}

func (d *dataScript) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dataScript) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["id"] = d.ID
	d.fieldMap["name"] = d.Name
	d.fieldMap["device_config_id"] = d.DeviceConfigID
	d.fieldMap["enable_flag"] = d.EnableFlag
	d.fieldMap["content"] = d.Content
	d.fieldMap["script_type"] = d.ScriptType
	d.fieldMap["last_analog_input"] = d.LastAnalogInput
	d.fieldMap["description"] = d.Description
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["remark"] = d.Remark
}

func (d dataScript) clone(db *gorm.DB) dataScript {
	d.dataScriptDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dataScript) replaceDB(db *gorm.DB) dataScript {
	d.dataScriptDo.ReplaceDB(db)
	return d
}

type dataScriptDo struct{ gen.DO }

type IDataScriptDo interface {
	gen.SubQuery
	Debug() IDataScriptDo
	WithContext(ctx context.Context) IDataScriptDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDataScriptDo
	WriteDB() IDataScriptDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDataScriptDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDataScriptDo
	Not(conds ...gen.Condition) IDataScriptDo
	Or(conds ...gen.Condition) IDataScriptDo
	Select(conds ...field.Expr) IDataScriptDo
	Where(conds ...gen.Condition) IDataScriptDo
	Order(conds ...field.Expr) IDataScriptDo
	Distinct(cols ...field.Expr) IDataScriptDo
	Omit(cols ...field.Expr) IDataScriptDo
	Join(table schema.Tabler, on ...field.Expr) IDataScriptDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDataScriptDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDataScriptDo
	Group(cols ...field.Expr) IDataScriptDo
	Having(conds ...gen.Condition) IDataScriptDo
	Limit(limit int) IDataScriptDo
	Offset(offset int) IDataScriptDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDataScriptDo
	Unscoped() IDataScriptDo
	Create(values ...*model.DataScript) error
	CreateInBatches(values []*model.DataScript, batchSize int) error
	Save(values ...*model.DataScript) error
	First() (*model.DataScript, error)
	Take() (*model.DataScript, error)
	Last() (*model.DataScript, error)
	Find() ([]*model.DataScript, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DataScript, err error)
	FindInBatches(result *[]*model.DataScript, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DataScript) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDataScriptDo
	Assign(attrs ...field.AssignExpr) IDataScriptDo
	Joins(fields ...field.RelationField) IDataScriptDo
	Preload(fields ...field.RelationField) IDataScriptDo
	FirstOrInit() (*model.DataScript, error)
	FirstOrCreate() (*model.DataScript, error)
	FindByPage(offset int, limit int) (result []*model.DataScript, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDataScriptDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dataScriptDo) Debug() IDataScriptDo {
	return d.withDO(d.DO.Debug())
}

func (d dataScriptDo) WithContext(ctx context.Context) IDataScriptDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dataScriptDo) ReadDB() IDataScriptDo {
	return d.Clauses(dbresolver.Read)
}

func (d dataScriptDo) WriteDB() IDataScriptDo {
	return d.Clauses(dbresolver.Write)
}

func (d dataScriptDo) Session(config *gorm.Session) IDataScriptDo {
	return d.withDO(d.DO.Session(config))
}

func (d dataScriptDo) Clauses(conds ...clause.Expression) IDataScriptDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dataScriptDo) Returning(value interface{}, columns ...string) IDataScriptDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dataScriptDo) Not(conds ...gen.Condition) IDataScriptDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dataScriptDo) Or(conds ...gen.Condition) IDataScriptDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dataScriptDo) Select(conds ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dataScriptDo) Where(conds ...gen.Condition) IDataScriptDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dataScriptDo) Order(conds ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dataScriptDo) Distinct(cols ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dataScriptDo) Omit(cols ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dataScriptDo) Join(table schema.Tabler, on ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dataScriptDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dataScriptDo) RightJoin(table schema.Tabler, on ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dataScriptDo) Group(cols ...field.Expr) IDataScriptDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dataScriptDo) Having(conds ...gen.Condition) IDataScriptDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dataScriptDo) Limit(limit int) IDataScriptDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dataScriptDo) Offset(offset int) IDataScriptDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dataScriptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDataScriptDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dataScriptDo) Unscoped() IDataScriptDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dataScriptDo) Create(values ...*model.DataScript) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dataScriptDo) CreateInBatches(values []*model.DataScript, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dataScriptDo) Save(values ...*model.DataScript) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dataScriptDo) First() (*model.DataScript, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DataScript), nil
	}
}

func (d dataScriptDo) Take() (*model.DataScript, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DataScript), nil
	}
}

func (d dataScriptDo) Last() (*model.DataScript, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DataScript), nil
	}
}

func (d dataScriptDo) Find() ([]*model.DataScript, error) {
	result, err := d.DO.Find()
	return result.([]*model.DataScript), err
}

func (d dataScriptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DataScript, err error) {
	buf := make([]*model.DataScript, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dataScriptDo) FindInBatches(result *[]*model.DataScript, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dataScriptDo) Attrs(attrs ...field.AssignExpr) IDataScriptDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dataScriptDo) Assign(attrs ...field.AssignExpr) IDataScriptDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dataScriptDo) Joins(fields ...field.RelationField) IDataScriptDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dataScriptDo) Preload(fields ...field.RelationField) IDataScriptDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dataScriptDo) FirstOrInit() (*model.DataScript, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DataScript), nil
	}
}

func (d dataScriptDo) FirstOrCreate() (*model.DataScript, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DataScript), nil
	}
}

func (d dataScriptDo) FindByPage(offset int, limit int) (result []*model.DataScript, count int64, err error) {
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

func (d dataScriptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dataScriptDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dataScriptDo) Delete(models ...*model.DataScript) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dataScriptDo) withDO(do gen.Dao) *dataScriptDo {
	d.DO = *do.(*gen.DO)
	return d
}
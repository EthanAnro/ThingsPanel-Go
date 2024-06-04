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

func newVisFile(db *gorm.DB, opts ...gen.DOOption) visFile {
	_visFile := visFile{}

	_visFile.visFileDo.UseDB(db, opts...)
	_visFile.visFileDo.UseModel(&model.VisFile{})

	tableName := _visFile.visFileDo.TableName()
	_visFile.ALL = field.NewAsterisk(tableName)
	_visFile.ID = field.NewString(tableName, "id")
	_visFile.VisPluginID = field.NewString(tableName, "vis_plugin_id")
	_visFile.FileName = field.NewString(tableName, "file_name")
	_visFile.FileURL = field.NewString(tableName, "file_url")
	_visFile.FileSize = field.NewString(tableName, "file_size")
	_visFile.CreateAt = field.NewInt64(tableName, "create_at")
	_visFile.Remark = field.NewString(tableName, "remark")

	_visFile.fillFieldMap()

	return _visFile
}

type visFile struct {
	visFileDo

	ALL         field.Asterisk
	ID          field.String
	VisPluginID field.String // 可视化插件id
	FileName    field.String // 名称
	FileURL     field.String // url地址
	FileSize    field.String // 文件大小
	CreateAt    field.Int64
	Remark      field.String

	fieldMap map[string]field.Expr
}

func (v visFile) Table(newTableName string) *visFile {
	v.visFileDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v visFile) As(alias string) *visFile {
	v.visFileDo.DO = *(v.visFileDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *visFile) updateTableName(table string) *visFile {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewString(table, "id")
	v.VisPluginID = field.NewString(table, "vis_plugin_id")
	v.FileName = field.NewString(table, "file_name")
	v.FileURL = field.NewString(table, "file_url")
	v.FileSize = field.NewString(table, "file_size")
	v.CreateAt = field.NewInt64(table, "create_at")
	v.Remark = field.NewString(table, "remark")

	v.fillFieldMap()

	return v
}

func (v *visFile) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *visFile) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 7)
	v.fieldMap["id"] = v.ID
	v.fieldMap["vis_plugin_id"] = v.VisPluginID
	v.fieldMap["file_name"] = v.FileName
	v.fieldMap["file_url"] = v.FileURL
	v.fieldMap["file_size"] = v.FileSize
	v.fieldMap["create_at"] = v.CreateAt
	v.fieldMap["remark"] = v.Remark
}

func (v visFile) clone(db *gorm.DB) visFile {
	v.visFileDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v visFile) replaceDB(db *gorm.DB) visFile {
	v.visFileDo.ReplaceDB(db)
	return v
}

type visFileDo struct{ gen.DO }

type IVisFileDo interface {
	gen.SubQuery
	Debug() IVisFileDo
	WithContext(ctx context.Context) IVisFileDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVisFileDo
	WriteDB() IVisFileDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVisFileDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVisFileDo
	Not(conds ...gen.Condition) IVisFileDo
	Or(conds ...gen.Condition) IVisFileDo
	Select(conds ...field.Expr) IVisFileDo
	Where(conds ...gen.Condition) IVisFileDo
	Order(conds ...field.Expr) IVisFileDo
	Distinct(cols ...field.Expr) IVisFileDo
	Omit(cols ...field.Expr) IVisFileDo
	Join(table schema.Tabler, on ...field.Expr) IVisFileDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVisFileDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVisFileDo
	Group(cols ...field.Expr) IVisFileDo
	Having(conds ...gen.Condition) IVisFileDo
	Limit(limit int) IVisFileDo
	Offset(offset int) IVisFileDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVisFileDo
	Unscoped() IVisFileDo
	Create(values ...*model.VisFile) error
	CreateInBatches(values []*model.VisFile, batchSize int) error
	Save(values ...*model.VisFile) error
	First() (*model.VisFile, error)
	Take() (*model.VisFile, error)
	Last() (*model.VisFile, error)
	Find() ([]*model.VisFile, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VisFile, err error)
	FindInBatches(result *[]*model.VisFile, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VisFile) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVisFileDo
	Assign(attrs ...field.AssignExpr) IVisFileDo
	Joins(fields ...field.RelationField) IVisFileDo
	Preload(fields ...field.RelationField) IVisFileDo
	FirstOrInit() (*model.VisFile, error)
	FirstOrCreate() (*model.VisFile, error)
	FindByPage(offset int, limit int) (result []*model.VisFile, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVisFileDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v visFileDo) Debug() IVisFileDo {
	return v.withDO(v.DO.Debug())
}

func (v visFileDo) WithContext(ctx context.Context) IVisFileDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v visFileDo) ReadDB() IVisFileDo {
	return v.Clauses(dbresolver.Read)
}

func (v visFileDo) WriteDB() IVisFileDo {
	return v.Clauses(dbresolver.Write)
}

func (v visFileDo) Session(config *gorm.Session) IVisFileDo {
	return v.withDO(v.DO.Session(config))
}

func (v visFileDo) Clauses(conds ...clause.Expression) IVisFileDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v visFileDo) Returning(value interface{}, columns ...string) IVisFileDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v visFileDo) Not(conds ...gen.Condition) IVisFileDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v visFileDo) Or(conds ...gen.Condition) IVisFileDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v visFileDo) Select(conds ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v visFileDo) Where(conds ...gen.Condition) IVisFileDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v visFileDo) Order(conds ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v visFileDo) Distinct(cols ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v visFileDo) Omit(cols ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v visFileDo) Join(table schema.Tabler, on ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v visFileDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v visFileDo) RightJoin(table schema.Tabler, on ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v visFileDo) Group(cols ...field.Expr) IVisFileDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v visFileDo) Having(conds ...gen.Condition) IVisFileDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v visFileDo) Limit(limit int) IVisFileDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v visFileDo) Offset(offset int) IVisFileDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v visFileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVisFileDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v visFileDo) Unscoped() IVisFileDo {
	return v.withDO(v.DO.Unscoped())
}

func (v visFileDo) Create(values ...*model.VisFile) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v visFileDo) CreateInBatches(values []*model.VisFile, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v visFileDo) Save(values ...*model.VisFile) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v visFileDo) First() (*model.VisFile, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VisFile), nil
	}
}

func (v visFileDo) Take() (*model.VisFile, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VisFile), nil
	}
}

func (v visFileDo) Last() (*model.VisFile, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VisFile), nil
	}
}

func (v visFileDo) Find() ([]*model.VisFile, error) {
	result, err := v.DO.Find()
	return result.([]*model.VisFile), err
}

func (v visFileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VisFile, err error) {
	buf := make([]*model.VisFile, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v visFileDo) FindInBatches(result *[]*model.VisFile, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v visFileDo) Attrs(attrs ...field.AssignExpr) IVisFileDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v visFileDo) Assign(attrs ...field.AssignExpr) IVisFileDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v visFileDo) Joins(fields ...field.RelationField) IVisFileDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v visFileDo) Preload(fields ...field.RelationField) IVisFileDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v visFileDo) FirstOrInit() (*model.VisFile, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VisFile), nil
	}
}

func (v visFileDo) FirstOrCreate() (*model.VisFile, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VisFile), nil
	}
}

func (v visFileDo) FindByPage(offset int, limit int) (result []*model.VisFile, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v visFileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v visFileDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v visFileDo) Delete(models ...*model.VisFile) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *visFileDo) withDO(do gen.Dao) *visFileDo {
	v.DO = *do.(*gen.DO)
	return v
}

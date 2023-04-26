// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"index_Demo/gen/orm/model"
)

func newLoginStatus(db *gorm.DB, opts ...gen.DOOption) loginStatus {
	_loginStatus := loginStatus{}

	_loginStatus.loginStatusDo.UseDB(db, opts...)
	_loginStatus.loginStatusDo.UseModel(&model.LoginStatus{})

	tableName := _loginStatus.loginStatusDo.TableName()
	_loginStatus.ALL = field.NewAsterisk(tableName)
	_loginStatus.ID = field.NewInt32(tableName, "id")
	_loginStatus.Token = field.NewString(tableName, "jwt")
	_loginStatus.UserID = field.NewInt32(tableName, "user_id")

	_loginStatus.fillFieldMap()

	return _loginStatus
}

type loginStatus struct {
	loginStatusDo

	ALL    field.Asterisk
	ID     field.Int32
	Token  field.String
	UserID field.Int32

	fieldMap map[string]field.Expr
}

func (l loginStatus) Table(newTableName string) *loginStatus {
	l.loginStatusDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l loginStatus) As(alias string) *loginStatus {
	l.loginStatusDo.DO = *(l.loginStatusDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *loginStatus) updateTableName(table string) *loginStatus {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.Token = field.NewString(table, "jwt")
	l.UserID = field.NewInt32(table, "user_id")

	l.fillFieldMap()

	return l
}

func (l *loginStatus) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *loginStatus) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 3)
	l.fieldMap["id"] = l.ID
	l.fieldMap["jwt"] = l.Token
	l.fieldMap["user_id"] = l.UserID
}

func (l loginStatus) clone(db *gorm.DB) loginStatus {
	l.loginStatusDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l loginStatus) replaceDB(db *gorm.DB) loginStatus {
	l.loginStatusDo.ReplaceDB(db)
	return l
}

type loginStatusDo struct{ gen.DO }

func (l loginStatusDo) Debug() *loginStatusDo {
	return l.withDO(l.DO.Debug())
}

func (l loginStatusDo) WithContext(ctx context.Context) *loginStatusDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loginStatusDo) ReadDB() *loginStatusDo {
	return l.Clauses(dbresolver.Read)
}

func (l loginStatusDo) WriteDB() *loginStatusDo {
	return l.Clauses(dbresolver.Write)
}

func (l loginStatusDo) Session(config *gorm.Session) *loginStatusDo {
	return l.withDO(l.DO.Session(config))
}

func (l loginStatusDo) Clauses(conds ...clause.Expression) *loginStatusDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loginStatusDo) Returning(value interface{}, columns ...string) *loginStatusDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loginStatusDo) Not(conds ...gen.Condition) *loginStatusDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loginStatusDo) Or(conds ...gen.Condition) *loginStatusDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loginStatusDo) Select(conds ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loginStatusDo) Where(conds ...gen.Condition) *loginStatusDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loginStatusDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *loginStatusDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l loginStatusDo) Order(conds ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loginStatusDo) Distinct(cols ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loginStatusDo) Omit(cols ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loginStatusDo) Join(table schema.Tabler, on ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loginStatusDo) LeftJoin(table schema.Tabler, on ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loginStatusDo) RightJoin(table schema.Tabler, on ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loginStatusDo) Group(cols ...field.Expr) *loginStatusDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loginStatusDo) Having(conds ...gen.Condition) *loginStatusDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loginStatusDo) Limit(limit int) *loginStatusDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loginStatusDo) Offset(offset int) *loginStatusDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loginStatusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *loginStatusDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loginStatusDo) Unscoped() *loginStatusDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loginStatusDo) Create(values ...*model.LoginStatus) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loginStatusDo) CreateInBatches(values []*model.LoginStatus, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loginStatusDo) Save(values ...*model.LoginStatus) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loginStatusDo) First() (*model.LoginStatus, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginStatus), nil
	}
}

func (l loginStatusDo) Take() (*model.LoginStatus, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginStatus), nil
	}
}

func (l loginStatusDo) Last() (*model.LoginStatus, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginStatus), nil
	}
}

func (l loginStatusDo) Find() ([]*model.LoginStatus, error) {
	result, err := l.DO.Find()
	return result.([]*model.LoginStatus), err
}

func (l loginStatusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LoginStatus, err error) {
	buf := make([]*model.LoginStatus, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loginStatusDo) FindInBatches(result *[]*model.LoginStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loginStatusDo) Attrs(attrs ...field.AssignExpr) *loginStatusDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loginStatusDo) Assign(attrs ...field.AssignExpr) *loginStatusDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loginStatusDo) Joins(fields ...field.RelationField) *loginStatusDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loginStatusDo) Preload(fields ...field.RelationField) *loginStatusDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loginStatusDo) FirstOrInit() (*model.LoginStatus, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginStatus), nil
	}
}

func (l loginStatusDo) FirstOrCreate() (*model.LoginStatus, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginStatus), nil
	}
}

func (l loginStatusDo) FindByPage(offset int, limit int) (result []*model.LoginStatus, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l loginStatusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loginStatusDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loginStatusDo) Delete(models ...*model.LoginStatus) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loginStatusDo) withDO(do gen.Dao) *loginStatusDo {
	l.DO = *do.(*gen.DO)
	return l
}
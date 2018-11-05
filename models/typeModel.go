package models

import (
	"math"
	"time"

	"../utils"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//Type ..
type Type struct {
	ID           string    `json:"id"  db:"id"  `
	Name         string    `json:"name"  db:"name"  `
	NamePinyin1  string    `json:"namePinyin1"  db:"namePinyin1"  `
	NamePinyin2  string    `json:"namePinyin2" db:"namePinyin2"`
	CreateTime   time.Time `json:"createTime"  db:"createTime"  `
	CreateUserID string    `json:"createUserId" db:"createUserId"`
}

//TypeMapper ..
type TypeMapper struct {
	BaseMapper
}

//GetTypeMapper ..
func GetTypeMapper(db string) TypeMapper {
	var mp TypeMapper
	mp.TableName = "noteType"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m TypeMapper) Get(tx *sqlx.Tx, whereStr string, args ...interface{}) (r interface{}) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	sqlStr := "select * from " + m.TableName + whereStr
	item := &User{}
	r = m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m TypeMapper) GetList(pageIndex int, rowsInPage int, sort string, whereStr string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*User, 0)
	sqlStr := "select * from " + m.TableName
	if sort != "" {
		sqlStr += " order by " + sort
	}
	if whereStr != "" {
		sqlStr += whereStr
	}
	err := m.getItems(nil, pageIndex, rowsInPage, &item, sqlStr)
	utils.CheckErr(err)
	total := m.GetCount(nil, "")
	var pageCount int
	if rowsInPage != 0 {
		pageCountF := float64(total) / float64(rowsInPage)
		pageCount = int(math.Ceil(pageCountF))
	} else {
		pageCount = 0
	}
	r = PagingData(pageIndex, rowsInPage, pageCount, total, item)
	return
}

//Insert ..
func (m TypeMapper) Insert(tx *sqlx.Tx, item *Type) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,name,namePinyin1,namePinyin2,createTime,createUserId) values (?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.Name, item.NamePinyin1, item.NamePinyin2, item.CreateTime, item.CreateUserID}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m TypeMapper) Update(tx *sqlx.Tx, item *Type) int {
	sqlStr := "update " + m.TableName + " set id=?,name=?,namePinYin1=?,namePinYin2=?,createTime=?,createUserId=? where id=? "
	var args = []interface{}{item.ID, item.Name, item.NamePinyin1, item.NamePinyin2, item.CreateTime, item.CreateUserID, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m TypeMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}

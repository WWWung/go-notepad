package models

import (
	"math"
	"time"

	"../utils"

	"github.com/ilibs/gosql"
	"github.com/jmoiron/sqlx"
)

//Note ..
type Note struct {
	ID           string    `json:"id"  db:"id"  `
	TypeID       string    `json:"typeId" db:"typeId"`
	Name         string    `json:"name"  db:"name"  `
	NamePinyin1  string    `json:"namePinyin1"  db:"namePinyin1"  `
	NamePinyin2  string    `json:"namePinyin2" db:"namePinyin2"`
	CreateTime   time.Time `json:"createTime"  db:"createTime"  `
	CreateUserID string    `json:"createUserId" db:"createUserId"`
	Tags         string    `json:"tags" db:"tags"`
	UpdateTime   time.Time `json:"updateTime" db:"updateTime"`
	HTMLContent  string    `json:"htmlContent" db:"htmlContent"`
	TextContent  string    `json:"textContent" db:"textContent"`
}

//NoteMapper ..
type NoteMapper struct {
	BaseMapper
}

//GetNoteMapper ..
func GetNoteMapper(db string) NoteMapper {
	var mp NoteMapper
	mp.TableName = "note"
	if db == "" {
		db = "default"
	}
	mp.DB = gosql.Use(db)
	return mp
}

//Get ..
func (m NoteMapper) Get(tx *sqlx.Tx, whereStr string, args ...interface{}) (r interface{}) {
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	sqlStr := "select * from " + m.TableName + whereStr
	item := &Note{}
	r = m.getItem(tx, item, sqlStr, args...)
	return r
}

//GetList ..
func (m NoteMapper) GetList(pageIndex int, rowsInPage int, sort string, whereStr string) (r interface{}) {
	//准备参数
	defer func() {
		if e := recover(); e != nil {
			r = nil
		}
	}()
	item := make([]*Note, 0)
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
func (m NoteMapper) Insert(tx *sqlx.Tx, item *Note) (int, int) {
	sqlStr := "insert into " + m.TableName + "(id,typeId,name,namePinyin1,namePinyin2,createTime,createUserId,tags,updateTime,htmlContent,textContent) values (?,?,?,?,?,?,?,?,?,?,?)"
	var args = []interface{}{item.ID, item.TypeID, item.Name, item.NamePinyin1, item.NamePinyin2, item.CreateTime, item.CreateUserID, item.Tags, item.UpdateTime, item.HTMLContent, item.TextContent}
	id, count := m.insertItem(tx, sqlStr, args...)
	return id, count
}

//Update ..
func (m NoteMapper) Update(tx *sqlx.Tx, item *Note) int {
	sqlStr := "update " + m.TableName + " set id=?,typeId=?,name=?,namePinYin1=?,namePinYin2=?,createTime=?,createUserId=?,tags=?,updateTime=?,htmlContent=?,textContent=? where id=? "
	var args = []interface{}{item.ID, item.TypeID, item.Name, item.NamePinyin1, item.NamePinyin2, item.CreateTime, item.CreateUserID, item.Tags, item.UpdateTime, item.HTMLContent, item.TextContent, item.ID}
	count := m.deleteOrUpdateItems(tx, sqlStr, args...)
	return count
}

//Delete ..
func (m NoteMapper) Delete(tx *sqlx.Tx, id interface{}) int {
	sqlStr := "delete from " + m.TableName + " where id = ? "
	var args = []interface{}{id}
	return m.deleteOrUpdateItems(tx, sqlStr, args...)
}

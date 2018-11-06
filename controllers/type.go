package controllers

import (
	"errors"
	"time"

	"../models"
	"../utils"
	"github.com/jmoiron/sqlx"
)

//TypeController ..
type TypeController struct {
	UserController
}

//Add ..
func (c *TypeController) Add() models.Type {
	item := models.Type{}
	c.parseItem(&item, true)
	mp := models.GetTypeMapper("")

	c.initData(&item)
	mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := utils.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()
		_, count := mp.Insert(tx, &item)
		if count == 0 {
			panic("增加失败")
		}
		return nil
	})

	return item
}

//Get ..
func (c *TypeController) Get() {

}

//GetList ..
func (c *TypeController) GetList() interface{} {
	c.r.ParseForm()
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getStringFromForm("sort")
	sortDir := c.getStringFromForm("sortDir")
	if sort != "" {
		sort += " " + sortDir + " "
	}
	mp := models.GetTypeMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort, "")
	return r
}

//Update ..
func (c *TypeController) Update() {

}

//Delete ..
func (c *TypeController) Delete() {

}

func (c *TypeController) initData(item *models.Type) {
	item.NamePinyin1, item.NamePinyin2 = utils.ToPinYin1(item.Name)
	user := c.getUserFromSession()
	item.CreateUserID = user.ID
	item.CreateTime = time.Now()
}

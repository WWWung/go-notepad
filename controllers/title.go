package controllers

import (
	"errors"

	"../models"
	"../utils"
	"github.com/jmoiron/sqlx"
)

//NoteController ..
type NoteController struct {
	UserController
}

//Add ..
func (c *NoteController) Add() models.Note {
	item := models.Note{}
	c.parseItem(&item, true)
	mp := models.GetNoteMapper("")

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
func (c *NoteController) Get() {

}

//GetList ..
func (c *NoteController) GetList() interface{} {
	pageIndex := c.getPageIndex()
	rowsInPage := c.getRowsInPage()
	sort := c.getStringFromForm("sort")
	sortDir := c.getStringFromForm("sortDir")
	if sort != "" {
		sort += " " + sortDir + " "
	}
	mp := models.GetNoteMapper("")
	r := mp.GetList(pageIndex, rowsInPage, sort, "")
	return r
}

//Update ..
func (c *NoteController) Update() {

}

//Delete ..
func (c *NoteController) Delete() {

}

func (c *NoteController) initData(item *models.Note) {

}

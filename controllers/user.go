package controllers

import (
	"errors"
	"fmt"
	"time"

	"../models"
	"../utils"
	"github.com/jmoiron/sqlx"
)

//UserController ..
type UserController struct {
	CaptchaController
}

//Add ..
func (c *UserController) Add() string {
	item := models.User{}
	c.parseItem(&item, true)
	mp := models.GetUserMapper("")

	if mp.GetCount(nil, " name=?", item.Name) > 0 {
		panic("账号已存在")
	}

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
			panic("注册失败")
		}
		return nil
	})
	c.store.Set("user", item)
	err := c.store.Save()
	utils.CheckErr(err)

	//	获取session
	s, ok := c.store.Get("user")
	fmt.Println("session is ", s, ok)

	return item.Name
}

//Login ..
func (c *UserController) Login() string {
	mp := models.GetUserMapper("")
	item := models.User{}
	c.parseItem(&item, false)
	user := mp.Get(nil, " where name=? ", item.Name)
	if user == nil {
		return "账号不存在"
	}
	u := user.(*models.User)
	if utils.Encrypt(item.Password) != u.PwMD5 {
		panic("账号或密码错误")
	}
	return u.Name
}

func (c *UserController) initData(item *models.User) {
	item.PwMD5 = utils.Encrypt(item.Password)
	item.CreateTime = time.Now()
	item.NamePinyin1, item.NamePinyin2 = utils.ToPinYin1(item.Name)
	item.LastLoginTime = time.Now()
	item.Power = 1
}

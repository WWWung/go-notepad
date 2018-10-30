package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"../utils"
)

//BaseControllers ..
type BaseControllers struct {
	data map[string]interface{}
	w    http.ResponseWriter
	r    *http.Request
	I    interface{}
}

//HandleHTTPRequest ..
func (c *BaseControllers) HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	c.w = w
	c.r = r
	c.data = make(map[string]interface{})
	c.crossDomin()
	c.r.ParseForm()
	c.apply()
}

func (c *BaseControllers) apply() {
	defer func() {
		if err := recover(); err != nil {
			c.failure(err)
		}
	}()
	methodName := c.getMethod()
	obj := reflect.ValueOf(c.I)
	params := make([]reflect.Value, 0)
	m := obj.MethodByName(methodName)
	if !m.IsValid() {
		panic("方法" + methodName + "不存在")
	}
	r := m.Call(params)
	if len(r) == 0 {
		c.success(nil)
	} else {
		c.success(r[0].Interface())
	}
}

func (c *BaseControllers) getContext() http.ResponseWriter {
	return c.w
}

func (c *BaseControllers) success(data interface{}) {
	d := map[string]interface{}{
		"code": 0,
		"data": data,
	}
	c.data["json"] = d
	c.json()
}

func (c *BaseControllers) json() {
	rsl := utils.ToJSON(c.data["json"])
	c.w.Write([]byte(rsl))
}

func (c *BaseControllers) failure(err interface{}) {
	c.data["json"] = map[string]interface{}{
		"code": 1,
		"data": err,
	}
	c.json()
}

func (c *BaseControllers) getMethod() string {
	v := c.getStringFromForm("m")
	if v == "" {
		panic("方法名不能为空")
	}
	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

func (c *BaseControllers) getStringFromForm(key string) string {
	vs := c.r.Form[key]
	if vs == nil || len(vs) == 0 || vs[0] == "" {
		vs = c.r.PostForm[key]
	}
	if vs == nil || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (c *BaseControllers) getID() string {
	return c.getStringFromForm("id")
}

func (c *BaseControllers) crossDomin() {
	c.w.Header().Set("Access-Control-Allow-Origin", "*")
}

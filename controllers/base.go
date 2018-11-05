package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"../utils"
	"github.com/go-session/session"
	uuid "github.com/satori/go.uuid"
)

//BaseControllers ..
type BaseControllers struct {
	data  map[string]interface{} //	储存要返回客户端的数据
	w     http.ResponseWriter
	r     *http.Request
	I     interface{}   //	储存子类(为了反射获取方法)
	input input         //	储存post数据
	store session.Store //	储存session
}

//post 数据格式
type input struct {
	M    string
	ID   string
	Data interface{}
}

//HandleHTTPRequest ..
func (c *BaseControllers) HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	c.w = w
	c.r = r
	fmt.Println(r.URL.Path, "===========")
	c.data = make(map[string]interface{})

	store, err := session.Start(context.Background(), c.w, c.r)
	utils.CheckErr(err)
	c.store = store

	//	请求的Content-type为application/json 所以post里取得的数据为json格式
	//	把这个数据解析成input类型并且存到BaseControllers.input里面
	con, _ := ioutil.ReadAll(r.Body)
	i := new(input)
	err = json.Unmarshal(con, i)
	utils.CheckErr(err)
	c.input = *i
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

func (c *BaseControllers) parseItem(item interface{}, setID bool) {
	jsonData := utils.ToJSON(c.input.Data)
	err := json.Unmarshal([]byte(jsonData), item)
	if setID {
		uid, err := uuid.NewV4()
		utils.CheckErr(err)
		id := uid.String()
		v := reflect.ValueOf(item).Elem()
		v.FieldByName("ID").Set(reflect.ValueOf(id))
	}
	utils.CheckErr(err)
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
	fmt.Println(err, "err")
	c.data["json"] = map[string]interface{}{
		"code": 1,
		"data": utils.InterfaceToStr(err),
	}
	c.json()
}

func (c *BaseControllers) getMethod() string {
	v := c.input.M
	if v == "" {
		panic("方法名不能为空")
	}
	f1 := v[:1]
	f2 := v[1:]
	return strings.ToUpper(f1) + f2 //方法名第1个字母要大写
}

func (c *BaseControllers) getStringFromForm(key string) string {
	c.r.ParseForm()
	vs := c.r.Form[key]
	if vs == nil || len(vs) == 0 || vs[0] == "" {
		vs = c.r.PostForm[key]
	}
	if vs == nil || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

//	从url参数里获取pageIndex
func (c *BaseControllers) getPageIndex() int {
	p := c.getStringFromForm("pageIndex")
	pageIndex, err := strconv.Atoi(p)
	utils.CheckErr(err)
	return pageIndex
}

//	从url参数里获取rowsInPage
func (c *BaseControllers) getRowsInPage() int {
	p := c.getStringFromForm("rowsInPage")
	rowsInPage, err := strconv.Atoi(p)
	utils.CheckErr(err)
	return rowsInPage
}

func (c *BaseControllers) getID() string {
	return c.input.ID
}

func (c *BaseControllers) crossDomin() {
	c.w.Header().Set("Access-Control-Allow-Origin", "*")
}

package controllers

//TestController ..
type TestController struct {
	BaseControllers
}

//Get ..
func (c *TestController) Get() interface{} {
	d := make(map[string]string)
	d["key"] = "key"
	return d
}

package controllers

import (
	"github.com/mojocn/base64Captcha"
)

//CaptchaController ..
type CaptchaController struct {
	BaseControllers
}

//CaptchaImg ..
type CaptchaImg struct {
	ID     string
	Base64 string
}

//Captcha ..
func (c *CaptchaController) Captcha() CaptchaImg {
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	img := CaptchaImg{idKeyD, base64stringD}
	return img
}

//Verfiy ..
func (c *CaptchaController) Verfiy(img CaptchaImg) bool {
	verifyResult := base64Captcha.VerifyCaptcha(img.ID, img.Base64)
	return verifyResult
}

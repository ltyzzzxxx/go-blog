package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/app/http/controllers/api/v1"
	"go-blog/app/requests"
	"go-blog/pkg/captcha"
	"go-blog/pkg/logger"
	"go-blog/pkg/response"
	"go-blog/pkg/verifycode"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}

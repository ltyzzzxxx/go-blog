package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/app/http/controllers/api/v1"
	"go-blog/pkg/captcha"
	"go-blog/pkg/logger"
	"go-blog/pkg/response"
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

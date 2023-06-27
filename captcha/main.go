package main

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/dchest/captcha"
	"time"
)

var (
	ErrNotFound = errors.New("captcha: id not found")
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)

func main() {
	h := server.Default()

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})

	globalStore := captcha.NewMemoryStore(CollectNum, Expiration)
	captcha.SetCustomStore(globalStore)
	h.GET("/captcha", func(ctx context.Context, c *app.RequestContext) {
		captchaId := captcha.New()
		// set a cookie
		cookie := protocol.AcquireCookie()
		cookie.SetKey("captchaId")
		cookie.SetValue(captchaId)
		cookie.SetExpire(time.Now().Add(3600 * time.Second))
		cookie.SetPath("/")
		cookie.SetHTTPOnly(true)
		cookie.SetSecure(false)
		c.Response.Header.SetCookie(cookie)
		captcha.WriteImage(c, captchaId, captcha.StdWidth, captcha.StdHeight)
	})

	h.POST("/verify", func(ctx context.Context, c *app.RequestContext) {
		captchaId := c.Cookie("captchaId")
		captchaNumber := c.PostForm("captcha")
		pass := captcha.VerifyString(string(captchaId), captchaNumber)
		if pass {
			c.String(consts.StatusOK, "Verify Captcha Succeed")
		} else {
			c.String(consts.StatusUnauthorized, "Verify Captcha Failed")
		}
	})

	h.Spin()
}

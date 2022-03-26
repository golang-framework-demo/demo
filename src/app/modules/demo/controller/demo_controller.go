// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"src/app/modules/demo/service"
	s "src/storage"
)

type DemoController struct {
	srv *service.DemoService
}

func NewDemoController() *DemoController {
	return &DemoController{
		srv: service.NewDemoService(),
	}
}

func (c *DemoController) TestDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, storage.Y{
		"demo":    "success",
		"service": c.srv.Demo(),
		"error": fmt.Sprintf(
			"%v: %v",
			s.Nd(s.KeyDemo10001),
			s.Ed(s.KeyDemo10001),
		),
	})
	ctx.Abort()

	return
}

func (c *DemoController) SetRedisDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.SetRedisDemo())
	ctx.Abort()

	return
}

func (c *DemoController) GetRedisDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.GetRedisDemo())
	ctx.Abort()

	return
}

func (c *DemoController) Demo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.GetDemo())
	ctx.Abort()

	return
}

func (c *DemoController) DemoAll(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.GetDemoAll())
	ctx.Abort()

	return
}

func (c *DemoController) DemoPaginator(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "1")

	ctx.JSON(storage.StatusOK, c.srv.GetDemoPaginator(cast.ToInt(page), cast.ToInt(size)))
	ctx.Abort()

	return
}

func (c *DemoController) DemoDetails(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.GetDemoDetails())
	ctx.Abort()

	return
}

func (c *DemoController) DemoD(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.GetDd())
	ctx.Abort()

	return
}

// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
	"src/app/modules/demo/service"
)

type CryptoController struct {
	srv *service.CryptoService
}

func NewCryptoController() *CryptoController {
	return &CryptoController{
		srv: service.NewCryptoService(),
	}
}

func (c *CryptoController) MD5(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.MD5("test success"))
	ctx.Abort()

	return
}

func (c *CryptoController) SHA1(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.SHA1("test success"))
	ctx.Abort()

	return
}

func (c *CryptoController) SHA256(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.SHA256("test success"))
	ctx.Abort()

	return
}

func (c *CryptoController) HMAC(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.srv.HMAC("test success"))
	ctx.Abort()

	return
}

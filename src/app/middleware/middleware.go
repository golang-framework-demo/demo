// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/src/middleware"
	"github.com/golang-framework/mvc/storage"
)

type (
	M struct {
		Parent *middleware.M
	}
)

func New() *M {
	return &M{
		Parent: middleware.New(),
	}
}

func NoRouter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(storage.StatusOK, "no router for golang-fw-demo")
		ctx.Abort()
		return
	}
}

func NoMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(storage.StatusOK, "no method for golang-fw-demo")
		ctx.Abort()
		return
	}
}

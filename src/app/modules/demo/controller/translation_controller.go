// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/modules/languages"
	"github.com/golang-framework/mvc/storage"
	s "src/storage"
)

type TranslationController struct {
}

func NewTranslationController() *TranslationController {
	return &TranslationController{}
}

func (c *TranslationController) T(ctx *gin.Context) {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	res.Res = &storage.Y{
		"1": languages.T("TEST_T_1"),
		"2": languages.T("TEST_T_2"),
	}

	ctx.JSON(storage.StatusOK, res)
	ctx.Abort()

	return
}

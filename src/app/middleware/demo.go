// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
)

type MDemo struct {
}

func NewMDemo() *MDemo {
	return &MDemo{}
}

func (_ *MDemo) InitializedLanguage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Query("ln") {
		case "cn":
			storage.CurrentLanguage = storage.CN
			break

		case "en":
			storage.CurrentLanguage = storage.EN
			break

		default:
			storage.CurrentLanguage = storage.CN
			break
		}

		ctx.Next()
	}
}

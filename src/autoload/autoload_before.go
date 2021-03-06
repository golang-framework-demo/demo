// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/routes"
	"src/app/middleware"
	"src/app/modules/demo/controller"
)

func (ad *autoload) before() {

}

func (ad *autoload) mvcInitializedRouter() (*routes.AHC, *routes.M) {
	var (
		demoMw      = middleware.NewMDemo()
		demo        = controller.NewDemoController()
		translation = controller.NewTranslationController()
		crypto      = controller.NewCryptoController()
	)

	return &routes.AHC{middleware.NoRouter(), middleware.NoMethod()},
		&routes.M{
			property.Instance.Get("route.Tag.Demo", "").(string): {
				Middleware: &routes.AHC{},
				Adapter: map[*routes.I]*routes.AHC{
					//http://127.0.0.1:8577/s_demo/demo/test_demo
					routes.AiGET(routes.DefaultRelativePath): {demo.TestDemo},

					// http://127.0.0.1:8577/s_demo/_demo_set
					routes.AiGET("_demo_set"): {demo.SetRedisDemo},
					// http://127.0.0.1:8577/s_demo/_demo_get
					routes.AiGET("_demo_get"): {demo.GetRedisDemo},
					// http://127.0.0.1:8577/s_demo/_demo
					routes.AiGET("_demo"): {demo.Demo},
					// http://127.0.0.1:8577/s_demo/_demo_all
					routes.AiGET("_demo_all"): {demo.DemoAll},
					// http://127.0.0.1:8577/s_demo/_demo_paginator?page=2&size=2
					routes.AiGET("_demo_paginator"): {demo.DemoPaginator},
					// http://127.0.0.1:8577/s_demo/_demo_details
					routes.AiGET("_demo_details"): {demo.DemoDetails},
					// http://127.0.0.1:8577/s_demo/_demo_d
					routes.AiGET("_demo_d"): {demo.DemoD},

					// http://127.0.0.1:8577/s_demo/_translation?language=cn
					routes.AiGET("_translation"): {
						demoMw.InitializedLanguage(),
						translation.T,
					},

					// http://127.0.0.1:8577/s_demo/_crypto_md5
					routes.AiGET("_crypto_md5"): {crypto.MD5},
					// http://127.0.0.1:8577/s_demo/_crypto_sha1
					routes.AiGET("_crypto_sha1"): {crypto.SHA1},
					// http://127.0.0.1:8577/s_demo/_crypto_sha256
					routes.AiGET("_crypto_sha256"): {crypto.SHA256},
					// http://127.0.0.1:8577/s_demo/_crypto_hmac
					routes.AiGET("_crypto_hmac"): {crypto.HMAC},
					// http://127.0.0.1:8577/s_demo/_crypto_encode
					routes.AiGET("_crypto_aes_encode"): {crypto.AesEncode},
					// http://127.0.0.1:8577/s_demo/_crypto_decode?aes_ciphertext=xxx
					routes.AiGET("_crypto_aes_decode"): {crypto.AesDecode},
				},
			},
		}
}

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
		demo = controller.NewDemoController()
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
					// http://127.0.0.1:8577/s_demo/_demo_details
					routes.AiGET("_demo_details"): {demo.DemoDetails},
					// http://127.0.0.1:8577/s_demo/_demo_d
					routes.AiGET("_demo_d"): {demo.DemoD},
				},
			},
		}
}

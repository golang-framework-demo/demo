// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/routes"
	"src/app/middleware"
	"src/app/services/demo/controller"
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
					routes.Ai("{_}"): {demo.TestDemo},
				},
			},
		}
}

// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"fmt"

	"github.com/golang-framework/mvc/src/components/caches/redis"
	"github.com/golang-framework/mvc/src/service"
	"github.com/golang-framework/mvc/storage"
	"src/app/db/models/demo"
	s "src/storage"
)

type DemoService struct {
	parent *service.Service
	demoDB *demo.DemoModel
}

func NewDemoService() *DemoService {
	return &DemoService{
		parent: service.New(),
		demoDB: demo.NewDemoModel(),
	}
}

func (srv *DemoService) Demo() string {
	return "demo_service"
}

func (srv *DemoService) SetRedisDemo() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	r := redis.New(0)
	r.SetPrefix("golang-framework-demo")

	_, errRedisSetDemo := r.Set("redis_demo", "test redis demo success!")
	if errRedisSetDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errRedisSetDemo.Error()

		return res
	}

	return res
}

func (srv *DemoService) GetRedisDemo() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	r := redis.New(0)
	r.SetPrefix("golang-framework-demo")

	o, errRedisGetDemo := r.Get("redis_demo")
	if errRedisGetDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errRedisGetDemo.Error()

		return res
	}

	res.Res = &storage.Y{
		"redis_demo": o,
	}

	return res
}

func (srv *DemoService) GetDemo() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	conditions := &storage.Conditions{}
	conditions.Query = "id=?"
	conditions.QueryArgs = []interface{}{1}
	conditions.Columns = []string{"id", "name"}

	d, errDemo := srv.demoDB.GetDemo(conditions)
	if errDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errDemo.Error()

		return res
	}

	res.Res = &storage.Y{
		"demo": d,
	}

	return res
}

func (srv DemoService) GetDemoAll() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	conditions := &storage.Conditions{}
	conditions.Columns = []string{"id", "name"}

	d, errDemo := srv.demoDB.GetDemoAll(conditions)
	if errDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errDemo.Error()

		return res
	}

	res.Res = &storage.Y{
		"demos": d,
	}

	return res
}
func (srv DemoService) GetDemoPaginator(page, size int) *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	conditions := &storage.Conditions{}
	conditions.Limit = size
	conditions.Start = (page - 1) * size

	d, errDemo := srv.demoDB.GetDemoAll(conditions)
	if errDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errDemo.Error()

		return res
	}

	fmt.Println(srv.demoDB.NumDemo())

	res.Res = &storage.Y{
		"demo":      d,
		"paginator": srv.parent.Paginator(page, 10, size),
	}

	return res
}

func (srv *DemoService) GetDemoDetails() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	conditions := &storage.Conditions{}
	conditions.Table = "demo"
	conditions.Field = []string{"demo.id", "demo.name", "demo_details.content"}
	conditions.Joins = []*storage.Join{
		{
			JoinOperator: "INNER",
			TableName:    "demo_details",
			Condition:    "demo.id=demo_details.demo_id",
		},
	}
	//conditions.Query = "demo.id=?"
	//conditions.QueryArgs = []interface{}{1}

	d, errDemo := srv.demoDB.GetDemoDetails(conditions)
	if errDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errDemo.Error()

		return res
	}

	res.Res = &storage.Y{
		"demo_details": d,
	}

	return res
}

func (srv *DemoService) GetDd() *storage.Tpl {
	res := storage.FwTpl(s.Ed(s.KeyDemo10001))

	conditions := &storage.Conditions{}
	conditions.Table = "demo"
	conditions.Field = []string{"demo.id", "demo.name", "d.content"}
	conditions.Joins = []*storage.Join{
		{
			JoinOperator: "INNER",
			TableName:    "d_demo",
			Condition:    "demo.id=d_demo.demo_id",
		},
		{
			JoinOperator: "INNER",
			TableName:    "d",
			Condition:    "d_id=d_demo.d_id",
		},
	}

	d, errDemo := srv.demoDB.GetDemoToD(conditions)
	if errDemo != nil {
		res.Status = storage.StatusUnknown
		res.Msg = errDemo.Error()

		return res
	}

	res.Res = &storage.Y{
		"demo_details": d,
	}

	return res
}

// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo

import (
	"github.com/golang-framework/mvc/src/models"
	"github.com/golang-framework/mvc/storage"
)

type DemoModel struct {
	mod *models.Model
}

func NewDemoModel() *DemoModel {
	return &DemoModel{
		mod: models.New(0),
	}
}

func (m *DemoModel) GetDemo(conditions *storage.Conditions) (*Demo, error) {
	conditions.Types = storage.SelectOne

	d := &Demo{}
	_, errDemo := m.mod.Select(conditions, d)

	return d, errDemo
}

func (m *DemoModel) GetDemoAll(conditions *storage.Conditions) ([]Demo, error) {
	conditions.Types = storage.SelectAll

	d := make([]Demo, 0)
	_, errDemo := m.mod.Select(conditions, &d)

	return d, errDemo
}

func (m *DemoModel) GetDemoDetails(conditions *storage.Conditions) ([]DemoToDetails, error) {
	conditions.Types = storage.SelectAll

	d := make([]DemoToDetails, 0)
	_, errDemoToDetails := m.mod.Select(conditions, &d)

	return d, errDemoToDetails
}

func (m *DemoModel) GetDemoToD(conditions *storage.Conditions) ([]DemoToD, error) {
	conditions.Types = storage.SelectAll

	d := make([]DemoToD, 0)
	_, errDemoToDemo := m.mod.Select(conditions, &d)

	return d, errDemoToDemo
}

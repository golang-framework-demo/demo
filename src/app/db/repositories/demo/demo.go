package demo

import (
	"time"
)

type Demo struct {
	Id         int        `xorm:"not null pk autoincr INT(11)" json:"id,omitempty"`
	Name       string     `xorm:"VARCHAR(255)" json:"name,omitempty"`
	Jointime   *time.Time `xorm:"DATETIME" json:"jointime,omitempty"`
	UpdateTime *time.Time `xorm:"updated TIMESTAMP" json:"update_time,omitempty"`
}

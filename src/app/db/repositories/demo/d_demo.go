package demo

type DDemo struct {
	Id     int `xorm:"not null pk INT(11)"`
	DId    int `xorm:"INT(11)"`
	DemoId int `xorm:"INT(11)"`
}

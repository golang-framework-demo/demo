package demo

type DemoDetails struct {
	Id      int    `xorm:"not null pk autoincr INT(11)" json:"id,omitempty"`
	DemoId  int    `xorm:"INT(11)" json:"demo_id,omitempty"`
	Content string `xorm:"VARCHAR(255)" json:"content,omitempty"`
}

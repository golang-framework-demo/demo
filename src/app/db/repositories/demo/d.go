package demo

type D struct {
	Id      int    `xorm:"not null pk INT(11)" json:"id,omitempty"`
	Content string `xorm:"VARCHAR(255)" json:"content,omitempty"`
}

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Guess struct {
	Id     int64
	Name   string `orm:"size(128)"`
	Age    int
	Option string
	Img    string
}

func init() {
	orm.RegisterModel(new(Guess))
}

// GetGuessById retrieves Guess by Id. Returns error if
// Id doesn't exist
func GetGuessById(id int64) (v Guess, err error) {
	o := orm.NewOrm()

	v = Guess{Id: id}
	err = o.Read(&v)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(v.Id, v.Name)
	}
	return v, err
}

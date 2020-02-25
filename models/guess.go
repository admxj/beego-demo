package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Guess struct {
	Id     int
	Name   string `orm:"size(128)"`
	Age    int
	Option string
	Img    string
	Answer string
}

func init() {
	orm.RegisterModel(new(Guess))
}

// GetGuessById retrieves Guess by Id. Returns error if
// Id doesn't exist
func GetGuessById(id int) (v Guess, err error) {
	o := orm.NewOrm()

	v = Guess{Id: id}
	err = o.Read(&v)

	return v, err
}

func Answer(sid int, answerkey string) bool {
	subject, err := GetGuessById(sid)

	if err != nil {
		return false
	}
	return strings.Compare(strings.ToUpper(answerkey), subject.Answer) == 0
}

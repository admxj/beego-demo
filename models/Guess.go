package models

import "github.com/astaxie/beego/orm"

type Guess struct {
	Id   int64
	Name string `orm:"size(128)"`
	Age  int
}

func init() {
	orm.RegisterModel(new(Guess))
}

// GetGuessById retrieves Guess by Id. Returns error if
// Id doesn't exist
func GetGuessById(id int64) (v *Guess, err error) {
	o := orm.NewOrm()
	v = &Guess{Id: id}
	if err = o.QueryTable(new(Guess)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

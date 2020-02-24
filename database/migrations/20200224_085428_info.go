package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Info_20200224_085428 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Info_20200224_085428{}
	m.Created = "20200224_085428"

	migration.Register("Info_20200224_085428", m)
}

// Run the migrations
func (m *Info_20200224_085428) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE info(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,`age` int(11) DEFAULT NULL)")
}

// Reverse the migrations
func (m *Info_20200224_085428) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `info`")
}

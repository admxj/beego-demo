package controllers

import (
	"github.com/astaxie/beego"
	"immooc/models"
)

//  InfoController operations for Info
type GuessController struct {
	beego.Controller
}

func (c *GuessController) Get() {
	var subject models.Guess
	err := func() {
		id, err := c.GetInt64("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		subject, err = models.GetGuessById(id)
	}
}

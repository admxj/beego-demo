package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/go-errors/errors"
	"immooc/models"
)

type GuessController struct {
	beego.Controller
}

func (c *GuessController) Get() {
	var guess models.Guess
	err := func() error {
		id, err := c.GetInt("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		guess, err = models.GetGuessById(id)
		if err != nil {
			return errors.New("not exist")
		}

		return nil
	}()
	if err != nil {
		c.Ctx.WriteString("wrong params")
	}
	var option map[string]string
	if err = json.Unmarshal([]byte(guess.Option), &option); err != nil {
		c.Ctx.WriteString("wrong response")
	}
	c.Data["ID"] = guess.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + guess.Img
	c.TplName = "guess.tpl"
}

func (c *GuessController) Post() {

	var guess models.Guess

	err := func() error {
		id, err := c.GetInt("id")
		beego.Info(id)
		if err != nil {
			id = 1
		}
		guess, err = models.GetGuessById(id)
		if err != nil {
			return errors.New("guess not exists")
		}
		return nil
	}()
	if err != nil {
		//c.Ctx.WriteString("wrong response")
	}
	answer := c.GetString("key")
	right := models.Answer(guess.Id, answer)

	c.Data["Right"] = right
	c.Data["Next"] = guess.Id + 1
	c.Data["ID"] = guess.Id
	c.TplName = "guess.tpl"
}

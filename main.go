package main


import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"gitlab.com/NicoNex/echotron"
)


type bot struct {
	chatId int64
	echotron.Engine
}


const BOTNAME = "Willy il Robot"


func NewBot(engine echotron.Engine, chatId int64) echotron.Bot {
	return &bot{
		chatId,
		engine,
	}
}


func (b bot) sendInitialMessage() {
	msg := `*Willy il Robot*

Digita una query e Willy cercherà nello storico di TNTVillage.
`

	b.SendMessageOptions(msg, b.chatId, echotron.PARSE_MARKDOWN)
}


func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.sendInitialMessage()
    }
}


func main() {
	home := os.Getenv("HOME")
	filepath := fmt.Sprintf("%s/.config/willyilrobot/TOKEN", home)

	token, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	echotron.RunDispatcher(string(token), NewBot)
}

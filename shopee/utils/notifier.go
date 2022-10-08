package utils

import (
	"errors"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v2"
	"shopee.rd/config"
)

var glb_notifier Notifier

type Notifier struct {
	Notifier_config config.Notifier_config
	Bot             *tgbotapi.BotAPI
	MaxRetryBackOff int
	Max_failure     int
}

func (notifier *Notifier) Init(config config.Notifier_config, maxRetryBackOff int, max_failure int) error {
	notifier.Notifier_config = config
	notifier.MaxRetryBackOff = maxRetryBackOff
	notifier.Max_failure = max_failure
	i := 0
	for i < notifier.Max_failure {
		if err := notifier.Connect(); err == nil {
			log.Println("Notifier connected")
			return nil
		} else {
			log.Println("Can't connect to notifier. Time ", i+1)
		}
		time.Sleep(time.Second * time.Duration(notifier.MaxRetryBackOff))
	}
	return errors.New("Connect to Notifier failed")
}
func (notifier *Notifier) Connect() error {
	log.Println(notifier.Notifier_config.Config.Bot_Api)
	bot, err := tgbotapi.NewBotAPI(notifier.Notifier_config.Config.Bot_Api)

	notifier.Bot = bot
	return err
}

func (notifier *Notifier) SendMessage(msg_str string) error {
	msg := tgbotapi.NewMessage(notifier.Notifier_config.Config.Bot_Chanel, msg_str)
	_, err := notifier.Bot.Send(msg)
	return err
}

func init() {
	if glb_notifier.Bot == nil {
		config_notifier := config.Notifier_config{}
		file, err := os.Open("/home/ubuntu/Documents/vscode/shopee/config/env-dev.yaml")
		if err != nil {
			log.Println("Can't open file shopee/config/env-dev.yaml")
			return
		}
		defer file.Close()
		d := yaml.NewDecoder(file)
		if err = d.Decode(&config_notifier); err != nil {
			log.Println("Can't decode to object file shopee/config/env-dev.yaml")
			return
		}
		if err = glb_notifier.Init(config_notifier, 15, 5); err != nil {
			log.Println(err)
			return
		} else {
			log.Println("Connected successfully")
			return
		}
	}
}
func Get_notifier() Notifier {
	return glb_notifier
}

package telegram

import (
	"log"

	"github.com/MerBasNik/telegram-bot/pkg/config"
	"github.com/MerBasNik/telegram-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot 		 *tgbotapi.BotAPI
	pocketClient *pocket.Client
	tokenRepository repository.TokenRepository
	redirectUrl string

	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, tr repository.TokenRepository, redirectUrl string, messages config.Messages) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, tokenRepository: tr, redirectUrl: redirectUrl, messages: messages}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel)  {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		} 
		
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message .Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}


func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}

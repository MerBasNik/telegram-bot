package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

//msg.Text = "Ты не авторизирован! используй команду /start"
//msg.Text = "Это невалидная ссылка"
//msg.Text = "Увы, не удалось сохранить вашу ссылку. Попробуй ещё раз позже"

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)

	switch err {
	case errUnauthorized:
		msg.Text = b.messages.Start
		b.bot.Send(msg)
	case errInvalidURL:
		msg.Text = b.messages.InvalidURL
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.messages.UnableToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
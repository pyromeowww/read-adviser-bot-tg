package telegram

import "read-adviser-bot/clients/telegram"

type Processor struct {
	tg *telegram.Client // Телеграм клиент
	offset int // Внутренний параметр offset
	// storage // storage нужен для того, чтобы куда-нибудь сохранять ссылки
}

// New - будет создавать процессор 
func New(client *telegram.Client)
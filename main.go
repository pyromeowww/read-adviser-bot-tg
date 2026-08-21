package main

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)


// mustToken - создаёт токен пакетом flag
func mustToken() string {
	token := flag.String("bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" { // если токен пуст, вызываем log.Fatal
		log.Fatal("token is not specified")
	}

	return *token
}

func main() {
	tgClient := telegram.New(tgBotHost, mustToken()) // Телеграм клиент

	/* fetcher и processor будут общаться с API телеграмма
	fetcher будет отправлять запрос чтобы получать новые события,
	а процессор после обработки сам будет отправлять новые сообщения
	*/

	// fetcher = fetcher.New(tgClient) Создаём fetcher

	// processor = processor.New(tgClient) Создаём processor

	//consumer получает и обрабатывает события
	// consumer.Start(fetcher, processor) для получение используется fetcher, а для обработки processor
}

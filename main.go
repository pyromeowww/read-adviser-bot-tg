package main

func main() {

	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	/* fetcher и processor будут общаться с API телеграмма
	fetcher будет отправлять запрос чтобы получать новые события,
	а процессор после обработки сам будет отправлять новые сообщения
	*/

	// fetcher = fetcher.New(tgClient) Создаём fetcher

	// processor = processor.New(tgClient) Создаём processor

	//consumer получает и обрабатывает события
	// consumer.Start(fetcher, processor) для получение используется fetcher, а для обработки processor
}

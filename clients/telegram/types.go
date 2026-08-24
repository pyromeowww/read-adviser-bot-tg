package telegram

// Тут определяем все типы с которыми будет работать наш клиент

type Chat struct {
	ID int `json:"id"`
}

type From struct {
	Username string `json:"username"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type UpdatesResponse struct {
	Ok     bool     `json: "ok"`
	Result []Update `json: "result"`
}

type Update struct {
	ID      int    `json: "update_id"`
	Message *IncomingMessage `json: "message"`
}

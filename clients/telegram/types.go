package telegram

// Тут определяем все типы с которыми будет работать наш клиент

type UpdatesResponse struct {
	Ok     bool     `json: "ok"`
	Result []Update `json: "result"`
}

type Update struct {
	ID      int    `json: "update_id"`
	Message string `json: "message"`
}

package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	e "read-adviser-bot/lib"
	"strconv"
	"time"
)

type Client struct {
	host     string      // host API сервиса telegram
	basePath string      // basePath тот префикс с которого начинаются все запросы
	client   http.Client // httpClient для того, чтобы не создавать его для каждого запроса отдельно
}

const (
	getUpdates        = "getUpdates"
	sendMessageMethod = "SendMessage"
)

func newBasePath(token string) string {
	return "bot" + token
}

// Отправка сообщений
func (c *Client) SendMessage(chatID int, text string) error {
	// подгатавливаем параматры запроса
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID)) // Добавляем указанный параметр к запросу
	q.Add("text", text)                    // Добавляем указанный параметр к запросу

	// Выполняем запрос, без тела ответа
	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

// Получение новых сообщений
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}                     // Формируем параметры запроса
	q.Add("offset", strconv.Itoa(offset)) // Добавляем указанный параметр к запросу
	q.Add("limit", strconv.Itoa(limit))   // Добавляем указанный параметр к запросу

	// Отправляем запрос
	data, err := c.doRequest(getUpdates, q)
	if err != nil {
		return nil, err
	}

	// Результат парсинга будет сохранять в переменную res
	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

// New - создаёт клиент
func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// Отправка запросов. do request <- getUpdates
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https", // протокол
		Host:   c.host,  // хост
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil) // Подгатавливаем запрос
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode() // Передаём параметры запроса
	resp, err := c.client.Do(req)     // Отправляем запрос
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }() // Закрываем тело ответа в конце
	body, err := io.ReadAll(resp.Body)       // Получаем содержимое
	if err != nil {
		return nil, err
	}

	return body, nil
}

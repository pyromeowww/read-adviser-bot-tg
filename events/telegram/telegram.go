package telegram

import (
	"errors"
	"read-adviser-bot/clients/telegram"
	"read-adviser-bot/events"
	e "read-adviser-bot/lib"
	"read-adviser-bot/storage"
)

type Meta struct {
	ChatID   int
	Username string
}

type Processor struct {
	tg      *telegram.Client // Телеграм клиент
	offset  int              // Внутренний параметр offset
	storage storage.Storage  // storage нужен для того, чтобы куда-нибудь сохранять ссылки
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

// New - будет создавать процессор
func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}
	return res, nil
}

// Метод Fetch
func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	// Получаем все апдейты с помощью клиента
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	// Память под результат
	res := make([]events.Event, 0, len(updates))

	// Обходим все аптейды и преобразуем в ивенты
	for _, u := range updates {
		res = append(res, event(u))
	}

	// Обновляем значение внутреннего поля offset
	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

// Метод process будет выполнять различные действия в зависимости от типа event
func (p Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	// Получаем Meta
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil
}

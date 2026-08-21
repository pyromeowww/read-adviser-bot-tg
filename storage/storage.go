package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
	e "read-adviser-bot/lib"
)

type Storage interface {
	Save(p *Page) error // Сохраняет страницу
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error) // Проверяет существует ли страница
}

// Основной пакет данных с которым будет работать Storage
type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha256.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	// Преобразуем байты в текст
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

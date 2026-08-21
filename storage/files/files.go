package files

import (
	"os"
	"path/filepath"
	e "read-adviser-bot/lib"
	"read-adviser-bot/storage"
)

type Storage struct {
	basePath string // Хранит информацию о том в какой папке мы будем всё хранить
}

const defaultPerm = 0774

// New - создаёт хранилище
func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	// Определяем куда мы будем сохранять наш файл
	filepath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filepath, defaultPerm); err != nil {
		return err
	}

	// Определяемся с названием файла

}

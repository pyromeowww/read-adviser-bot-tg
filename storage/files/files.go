package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	e "read-adviser-bot/lib"
	"read-adviser-bot/storage"
)

type Storage struct {
	basePath string // Хранит информацию о том в какой папке мы будем всё хранить
}

const defaultPerm = 0774

var ErrNoSavedPages = errors.New("no saved pages")

// New - создаёт хранилище
func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

// Метод IsExists - будет говорить существует данная страница или нет
func (s Storage) IsExists(p *storage.Page) (bool, error) {

	fileName, err := filename(p)
	if err != nil {
		return false, e.Wrap("Can't remove file", err)
	}
	// Получаем путь до файла
	path := filepath.Join(s.basePath, p.UserName, fileName)

	// Проверяем существование
	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)
		return false, e.Wrap(msg, err)
	}

	return true, nil
}

// Метод PickRandom
func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("can't pick random page", err) }() // Определяем способо обработки ошибок

	// Получаем путь до директории с файлами
	path := filepath.Join(s.basePath, userName)

	// Получаем список файлов
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	// Получаем рандомное число от 0 до количества файлов
	n := rand.IntN(len(files))

	// Получаем случайный файл

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

// Метод Remove

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := filename(p)
	if err != nil {
		return e.Wrap("Can't remove file", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	msg := fmt.Sprintf("can't remove file %s", path)
	if err := os.Remove(path); err != nil {
		return e.Wrap(msg, err)
	}

	return nil
}

// Метод Save
func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save page", err) }() // Определяем способо обработки ошибок

	// Формируем путь куда будет сохраняться файл
	fPath := filepath.Join(s.basePath, page.UserName)

	// Создаём все нужные директории в пути
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	// Формируем имя файла
	fName, err := filename(page)
	if err != nil {
		return err
	}

	// Дописываем имя файла к пути
	fPath = filepath.Join(fPath, fName)

	// Создаём файл
	file, err := os.Create(fPath)

	defer func() { _ = file.Close() }()

	// Записываем в файл страницу в нужном формате
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode page", err)
	}
	// Закрываем файл
	defer func() { _ = file.Close() }()

	// Создаём переменную в которую файл будет декодирован
	var p storage.Page

	if err := gob.NewDecoder(file).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	return &p, nil
}

// filename - будет определять имя файла
func filename(p *storage.Page) (string, error) {
	return p.Hash() // Построение хеша
}

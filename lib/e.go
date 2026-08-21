package e

import "fmt"

// Wrap - занимается оборачиванием ошибок
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Если WrapIfErr получит нулевую ошибку, то вернёт нил
func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}

package task

import "errors"

var typeError = errors.New("type error")

type Config interface {
	Attr01() map[string]int
	Attr02() int
}

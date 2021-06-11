package task

import "log"

func Task02(i interface{}) (interface{}, error) {
	config, ok := i.(Config)
	if !ok {
		return nil, typeError
	}
	attr01 := config.Attr01()
	attr02 := config.Attr02()
	log.Printf("attr01 is %v, attr02 is %v", attr01, attr02)
	return nil, nil
}

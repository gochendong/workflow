package task

func Task01(i interface{}) (interface{}, error) {
	config, ok := i.(Config)
	if !ok {
		return nil, typeError
	}
	attr01 := config.Attr01()
	attr02 := config.Attr02()
	attr01["I am map"] = 1
	attr02 += 1 // 传递的数据需用map存储, 否则修改无效
	return config, nil
}

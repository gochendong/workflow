package cmd

import (
	"log"
	"os"
	"workflow/stream"
	"workflow/task"
)

const OSEXIT2 = 2

type config struct {
	attr01 map[string]int
	attr02 int
}

func Run() {
	log.SetPrefix("[work flow]")
	arg := parse()
	_, err := stream.NewStream().
		Next(task.Task01).
		Next(task.Task02).
		Go(arg)
	if err != nil {
		log.Println(err)
		os.Exit(OSEXIT2)
	}
	log.Println("SUCCESS")
}

func parse() *config {
	arg := &config{}
	arg.attr01 = make(map[string]int, 1) // map需要初始化
	return arg
}

func (c *config) Attr01() map[string]int {
	return c.attr01
}

func (c *config) Attr02() int {
	return c.attr02
}

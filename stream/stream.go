package stream

import (
	"errors"
)

/**
流式工作原理：
各个任务都过指针链表的方式组成一个任务链，这个任务链从第一个开始执行，直到最后一个
每一个任务节点执行完毕会将结果带入到下一级任务节点中。
每一个任务是一个Stream节点，每个任务节点都包含首节点和下一个任务节点的指针,
除了首节点，每个节都会设置一个回调函数的指针，用本节点的任务执行,
最后一个节点的nextStream为空,表示任务链结束。
**/

type CB func(interface{}) (interface{}, error)

type Stream struct {
	firstStream *Stream
	nextStream  *Stream
	cb          CB
}

/**
创建新的流
**/
func NewStream() *Stream {
	stream := &Stream{}
	stream.firstStream = stream
	return stream
}

/**
流结束
arg为流初始参数，初始参数放在End方法中是考虑到初始参数不需在任务链中传递
**/
func (this *Stream) Go(arg interface{}) (interface{}, error) {
	this.nextStream = nil
	if this.firstStream.nextStream != nil {
		return this.firstStream.nextStream.run(arg)
	} else {
		return nil, errors.New("Not found execute node.")
	}
}

func (this *Stream) run(arg interface{}) (interface{}, error) {
	result, err := this.cb(arg)
	if this.nextStream != nil && err == nil {
		return this.nextStream.run(result)
	} else {
		return result, err
	}
}

func (this *Stream) Next(cb CB) *Stream {
	this.nextStream = &Stream{}
	this.nextStream.firstStream = this.firstStream
	this.nextStream.cb = cb
	return this.nextStream
}

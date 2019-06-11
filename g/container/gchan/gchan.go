<<<<<<< HEAD
// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 优雅的Channel操作.
=======
// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gchan provides graceful channel for no panic operations.
//
// It's safe to call Chan.Push/Close functions repeatedly.
>>>>>>> upstream/master
package gchan

import (
    "errors"
<<<<<<< HEAD
    "gitee.com/johng/gf/g/container/gtype"
)

type Chan struct {
    list   chan interface{}
    closed *gtype.Bool
}

func New(limit int) *Chan {
    return &Chan {
        list   : make(chan interface{}, limit),
        closed : gtype.NewBool(),
    }
}

// 将数据压入队列
func (q *Chan) Push(v interface{}) error {
    if q.closed.Val() {
        return errors.New("closed")
    }
    q.list <- v
    return nil
}

// 先进先出地从队列取出一项数据，当没有数据可获取时，阻塞等待
func (q *Chan) Pop() interface{} {
    return <- q.list
}

// 关闭队列(通知所有通过Pop阻塞的协程退出)
func (q *Chan) Close() {
    if !q.closed.Val() {
        q.closed.Set(true)
        close(q.list)
    }
}

// 获取当前队列大小
func (q *Chan) Size() int {
    return len(q.list)
=======
    "github.com/gogf/gf/g/container/gtype"
)

// Graceful channel.
type Chan struct {
    channel chan interface{}
    closed  *gtype.Bool
}

// New creates a graceful channel with given <limit>.
func New(limit int) *Chan {
    return &Chan {
	    channel : make(chan interface{}, limit),
        closed  : gtype.NewBool(),
    }
}

// Push pushes <value> to channel.
// It is safe to be called repeatedly.
func (c *Chan) Push(value interface{}) error {
    if c.closed.Val() {
        return errors.New("channel is closed")
    }
    c.channel <- value
    return nil
}

// Pop pops value from channel.
// If there's no value in channel, it would block to wait.
// If the channel is closed, it will return a nil value immediately.
func (c *Chan) Pop() interface{} {
    return <- c.channel
}

// Close closes the channel.
// It is safe to be called repeatedly.
func (c *Chan) Close() {
    if !c.closed.Set(true) {
        close(c.channel)
    }
}

// See Len.
func (c *Chan) Size() int {
    return c.Len()
}

// Len returns the length of the channel.
func (c *Chan) Len() int {
	return len(c.channel)
>>>>>>> upstream/master
}
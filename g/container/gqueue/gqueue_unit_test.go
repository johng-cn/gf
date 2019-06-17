package gqueue_test

import (
	"github.com/gogf/gf/g/container/gqueue"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
)

func TestQueue_Len(t *testing.T) {
	maxs := 100
	for n := 10; n < maxs; n++ {
		q1 := gqueue.New(maxs)
		for i := 0; i < maxs; i++ {
			q1.Push(i)
		}
		gtest.Assert(q1.Len(), maxs)
	}
}

func TestQueue_Pop(t *testing.T) {
	q1 := gqueue.New()
	q1.Push(1)
	q1.Push(2)
	q1.Push(3)
	q1.Push(4)
	i1 := q1.Pop()
	gtest.Assert(i1, 1)

}

func TestQueue_Close(t *testing.T) {
	q1 := gqueue.New()
	q1.Push(1)
	q1.Push(2)
	gtest.Assert(q1.Len(), 2)
	q1.Close()
}
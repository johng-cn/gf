<<<<<<< HEAD
// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
=======
// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
>>>>>>> upstream/master

package gtype

import (
    "sync/atomic"
)

type Uint32 struct {
<<<<<<< HEAD
    val uint32
}

func NewUint32(value...uint32) *Uint32 {
    if len(value) > 0 {
        return &Uint32{val:value[0]}
=======
	value uint32
}

// NewUint32 returns a concurrent-safe object for uint32 type,
// with given initial value <value>.
func NewUint32(value...uint32) *Uint32 {
    if len(value) > 0 {
        return &Uint32{
        	value : value[0],
		}
>>>>>>> upstream/master
    }
    return &Uint32{}
}

<<<<<<< HEAD
func (t *Uint32)Set(value uint32) {
    atomic.StoreUint32(&t.val, value)
}

func (t *Uint32)Val() uint32 {
    return atomic.LoadUint32(&t.val)
}

func (t *Uint32)Add(delta uint32) uint32 {
    return atomic.AddUint32(&t.val, delta)
=======
// Clone clones and returns a new concurrent-safe object for uint32 type.
func (t *Uint32) Clone() *Uint32 {
    return NewUint32(t.Val())
}

// Set atomically stores <value> into t.value and returns the previous value of t.value.
func (t *Uint32) Set(value uint32) (old uint32) {
    return atomic.SwapUint32(&t.value, value)
}

// Val atomically loads t.value.
func (t *Uint32) Val() uint32 {
    return atomic.LoadUint32(&t.value)
}

// Add atomically adds <delta> to t.value and returns the new value.
func (t *Uint32) Add(delta uint32) (new uint32) {
    return atomic.AddUint32(&t.value, delta)
>>>>>>> upstream/master
}
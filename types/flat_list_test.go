package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatListLen(t *testing.T) {
	assert := assert.New(t)

	l := List(flatList{})
	assert.Equal(uint64(0), l.Len())
	l = l.Append(Bool(true))
	assert.Equal(uint64(1), l.Len())
	l = l.Append(Bool(false), Bool(false))
	assert.Equal(uint64(3), l.Len())
}

func TestFlatListGet(t *testing.T) {
	assert := assert.New(t)

	l := List(flatList{})
	l = l.Append(Int32(0), Int32(1), Int32(2))
	assert.Equal(Int32(0), l.Get(0))
	assert.Equal(Int32(1), l.Get(1))
	assert.Equal(Int32(2), l.Get(2))

	assert.Panics(func() {
		l.Get(3)
	})
}

func TestFlatListSlice(t *testing.T) {
	assert := assert.New(t)
	l1 := List(flatList{})
	l1 = l1.Append(Int32(0), Int32(1), Int32(2), Int32(3))
	l2 := l1.Slice(1, 3)
	assert.Equal(uint64(4), l1.Len())
	assert.Equal(uint64(2), l2.Len())
	assert.Equal(Int32(1), l2.Get(0))
	assert.Equal(Int32(2), l2.Get(1))

	l3 := l1.Slice(0, 0)
	assert.Equal(uint64(0), l3.Len())
	l3 = l1.Slice(1, 1)
	assert.Equal(uint64(0), l3.Len())
	l3 = l1.Slice(1, 2)
	assert.Equal(uint64(1), l3.Len())
	assert.Equal(Int32(1), l3.Get(0))
	l3 = l1.Slice(0, l1.Len())
	assert.True(l1.Equals(l3))

	assert.Panics(func() {
		l3 = l1.Slice(0, l1.Len()+1)
	})
}

func TestFlatListAppend(t *testing.T) {
	assert := assert.New(t)

	l0 := flatList{}
	l1 := l0.Append(Bool(false))
	assert.Equal(uint64(0), l0.Len())
	assert.Equal(uint64(1), l1.Len())
	assert.Equal(Bool(false), l1.Get(0))

	// Append(v1, v2)
	l2 := l1.Append(Bool(true), Bool(true))
	assert.Equal(uint64(3), l2.Len())
	assert.Equal(Bool(false), l2.Get(0))
	assert.True(NewList(Bool(true), Bool(true)).Equals(l2.Slice(1, l2.Len())))
	assert.Equal(uint64(1), l1.Len())
}
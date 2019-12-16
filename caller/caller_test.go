package caller

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)


// 指针可以调用非指针的函数
// 行为是按值来的，而不是按指针来，即不改变调用者
type A struct {
	i int
}

func (a A) Add(j int) A {
	fmt.Printf("addr add %p\n", &a)
	a.i += j
	return a
}

func Test_Modify(t *testing.T) {
	a := &A{i: 1}
	fmt.Printf("addr add %p\n", a)
	b := a.Add(1)
	fmt.Printf("addr add %p\n", &b)
	assert.Equal(t, 2, b.i)
	assert.Equal(t, 1, a.i)

}


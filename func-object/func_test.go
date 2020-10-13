package func_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义一个类型是函数类型, 但他可以像结构体一样使用
// 如下面, 给httpH 类型定义一个 Serve 成员函数
// 这样可以简单的实现接口适配

// 模拟http.HandlerFunc
type httpH func(string, int)

// Serve 实际即调用函数
func (f httpH) Serve(o string, i int) {
	f(o, i)
}

type httpHandler interface {
	Serve(o string, i int)
}

func Test_f1(t *testing.T) {
	var handler httpHandler = httpH(func(o string, i int) {
		fmt.Println("interface")
	})
	handler2 := func(o string, i int) { fmt.Println("func") }
	var handlerRaw httpH
	handlerRaw = handler2
	t.Log(reflect.TypeOf(handler), reflect.TypeOf(handler2), reflect.TypeOf(handlerRaw))
	assert.Nil(t, handlerRaw)
}

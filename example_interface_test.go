package msgpack_test

import (
	"fmt"
	"github.com/ansj/msgpack"
)

type Item struct {
	Foo string
}

func (i Item) Name() string{
	return i.Foo
}

type I interface {
	Name() string
}

func ExampleMarshal_Interface() {

	msgpack.RegisterType(Item{})

	var i I = &Item{Foo: "bar"}

	b, err := msgpack.MarshalInterface(i)
	if err != nil {
		panic(err)
	}

	var i2 I = &Item{}
	err = msgpack.UnmarshalInterface(b, i2)
	if err != nil {
		panic(err)
	}
	fmt.Println(i2.Name())
	// Output: bar
}
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ansj/msgpack"
)

type Item1 struct {
	Name1 string
}

type Item2 struct {
	Name2 string
}

type Items struct {
	Item1   I
	Arr     []I
	Arr1    int
	Item2   I
	ItemNil I
	ItemP   *I
	ArrP    []*I
	ArrP2   *[]I
	Name    string
	Mapf    map[string]I
	MapP    map[string]*I
}

func (this Items) GetName() string {
	return this.Name
}

func (this Item1) GetName() string {
	return this.Name1
}

func (this Item2) GetName() string {
	return this.Name2
}

type I interface {
	GetName() string
}

func main() {

	var tt I = Item1{}
	msgpack.Register(tt,nil,nil)

	msgpack.RegisterType(Item1{})
	msgpack.RegisterType(Item2{})

	arr := []I{
		Item1{"1",},
		Item1{"2",},
		Item2{"5",},
	}

	var p1 I = Item1{"p1",}
	var p2 I = Item1{"p1",}
	arrP := []*I{
		&p1,
		&p2,
	}

	mapf := map[string]I{
		"aaa": Item1{"1",},
		"bbb": Item2{"2",},
	}

	mapP := map[string]*I{
		"aaa": &p1,
		"bbb": &p2,
	}


	var item1 I = Item1{"itemPoint",}
	in := Items{
		Item1: Item1{"hello",},
		Item2: Item2{"ansj",},
		Name:  "ansj",
		Arr:   arr,
		ArrP:  arrP,
		ArrP2: &arr,
		Mapf:  mapf,
		MapP:  mapP,
		ItemP: &item1,
	}

	bytes, _ := msgpack.MarshalInterface(&in)

	out := Items{}
	fmt.Println(bytes)
	if err := msgpack.UnmarshalInterface(bytes, &out); err != nil {
		panic(err)
	}
	fmt.Println(out)

	marshal, _ := json.Marshal(out)
	fmt.Println(string(marshal))

}

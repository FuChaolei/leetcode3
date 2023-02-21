//go:build ignore

package main

import "fmt"

//"fmt"
//"container/list"
//"sort"
//"math/rand"
//"math"
//"strings"
type abstractFruiter interface {
	ShowName()
}
type Apple struct {
}

func (apple *Apple) ShowName() {
	fmt.Println("I'm apple")
}

type Banana struct {
}

func (banana *Banana) ShowName() {
	fmt.Println("I'm banana")
}

type abstractFactory struct {
}

func (factory *abstractFactory) Createfruit(name string) abstractFruiter {
	//var fruit abstractFactory
	if name == "apple" {
		return new(Apple)
	} else if name == "banana" {
		return new(Banana)
	}
	return nil
}
func main() {
	//res := get_res()
	factory := new(abstractFactory)
	fruit := factory.Createfruit("apple")
	fruit.ShowName()
	fruit = factory.Createfruit("banana")
	fruit.ShowName()
}

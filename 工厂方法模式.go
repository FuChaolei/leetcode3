//go:build ignore

package main

import "fmt"

//"fmt"
//"container/list"
//"sort"
//"math/rand"
//"math"
//"strings"
type Abstractfruit interface {
	showName()
}
type Apple struct {
}

func (apple *Apple) showName() {
	fmt.Println("我是苹果")
}

type Banana struct {
}

func (banana *Banana) showName() {
	fmt.Println("我是香蕉")
}

type Abstractfactory interface {
	Createfruit() Abstractfruit
}
type Applefactory struct {
}

func (applefactory *Applefactory) Createfruit() Abstractfruit {
	return new(Apple)
}

type Bananafactory struct {
}

func (bananafactory *Bananafactory) Createfruit() Abstractfruit {
	return new(Banana)
}

func main() {
	var factory Abstractfactory
	factory = new(Applefactory)
	var apple Abstractfruit
	apple = factory.Createfruit()
	apple.showName()
	factory = new(Bananafactory)
	var banana Abstractfruit
	banana = factory.Createfruit()
	banana.showName()
	//	fruit := factory.Createfruit()
}

//go:build ignore

package main

import "fmt"

//"fmt"
//"container/list"
//"sort"
//"math/rand"
//"math"
//"strings"
type Hungry struct{}

var hungry *Hungry = new(Hungry)

func getInstace() *Hungry {
	return hungry
}
func (it *Hungry) do() {
	fmt.Println("sdfasdf")
}

func main() {
	tmp := getInstace()
	tmp.do()
	tmp2 := getInstace()
	if tmp2 == tmp {
		fmt.Println("dsgfdg")
	}
}

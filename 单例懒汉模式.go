//go:build ignore

package main

import (
	"fmt"
	"sync"
)

//var Mutex sync.Mutex
var once sync.Once

type Lazy struct{}

var lazy *Lazy

func getInstance() *Lazy {
	// Mutex.Lock()
	// defer Mutex.Unlock()
	once.Do(func() {
		lazy = new(Lazy)
	})
	return lazy
}
func (someth *Lazy) dosome() {
	fmt.Println("come欧尼！")
}
func main() {
	tmp := getInstance()
	tmp.dosome()
}

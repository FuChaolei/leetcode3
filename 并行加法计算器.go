//go:build ignore

package main

import (
	"fmt"
	//	"runtime"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	res := 0
	add := func(s int) {
		for i := s; i < s+100; i++ {
			//fmt.Println(i)
			mu.Lock()
			res += i
			mu.Unlock()
		}
	}
	for i := 0; i < 1000; i++ {
		go add(i*100 + 1)
	}
	//fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
	fmt.Println(res)
	//res := get_res()
}

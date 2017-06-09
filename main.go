 package main

 import (
    "fmt"
	"time"
	"math/rand"
 )

var TestData         = [...]string{"This", "is", "a", "test."}
var TestData2         = [...]string{"Hope", "this ", "works", "good!"}

var commandQueue = make(chan string, 1) //Single Channel for handling this is all we need.
var QueueSleepInterval = struct{ min, max int }{1, 3}

func test(){
	var i int = 0
	for {
		wsi := QueueSleepInterval
		s := rand.Intn(wsi.max-wsi.min) + wsi.min
		time.Sleep(time.Duration(s) * time.Second)

		data := <-commandQueue
		fmt.Println(i, data)
		i++
	}
}

 func main() {	
	go test()
	for{
	for _, data := range TestData {
		commandQueue <- data
	}
	
	time.Sleep(5 * time.Second)
	
  	for _, data := range TestData2 {
		commandQueue <- data
	} 
	}
 }

// main.o
package main

import (
	"fmt"
	"library"
	"time"
)

func main() {
	grt := library.Greeting{}
	result := grt.GetGreetingByHours(time.Now())

	fmt.Println(result)
}

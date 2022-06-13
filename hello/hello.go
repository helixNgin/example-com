/**
    * @Description:
    * @Author: daniel(2022/6/13)
**/
package main

import (
	"example-com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("gg")
	fmt.Println(message)
}

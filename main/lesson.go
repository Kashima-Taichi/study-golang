package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("we will use time package", time.Now())
	fmt.Println(user.Current())
}

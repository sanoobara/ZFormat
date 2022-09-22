package main

import (
	"time"
	"fmt"
	"rm/internal/mover"
)

func main() {
	
	for {
		
		err := mover.MoveFiles("d:\\test1", "d:\\test2")
		if err != nil{
			fmt.Println(err)
		}
		time.Sleep(2*time.Second)
	}
	
}

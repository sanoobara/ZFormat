package main

import (
	"time"
	"fmt"
	"rm/internal/mover"
)

func main() {
	
	for {
		
		DirEntry, err := mover.GetListFiles("d:\\test1")
		err := mover.MoveFiles(DirEntry, "d:\\test2")
		if err != nil{
			fmt.Println(err)
		}
		time.Sleep(2*time.Second)
	}
	
}

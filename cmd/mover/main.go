package main

import (
	"fmt"
	"rm/internal/mover"
	"time"
)

func main() {

	for {

		Files, err := mover.GetListFiles("d:\\test1")
		err = mover.MoveFiles(Files, "d:\\test2")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(2 * time.Second)
	}

}

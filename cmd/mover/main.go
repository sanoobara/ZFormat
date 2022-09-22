package main

import (
	"fmt"
	"rm/internal/config"
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

		_, er := config.ParsHeader("config\\zheader.json")

		if er != nil {
			fmt.Println(er)
		}

	}

}

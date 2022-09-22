package main


import (
	"fmt"
	"rm/internal/config"

	
)

func main() {


		

		_, er := config.ParsHeader("config\\zheader.json")

		if er != nil {
			fmt.Println(er)
		}

	

}

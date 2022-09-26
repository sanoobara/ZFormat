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


	// _, er := config.ParsHeader("config\\zheader.json")

	// if er != nil {
	// 	fmt.Println(er)
	// }

	// //fmt.Println(filepath.Abs("\\\\sinto89\\SPO\\BULK_EXEC\\LOGS\\BulkExec.log"))

	// _, e:=config.MainConfig("config\\config.json")
	
	// if e != nil {
	// 	fmt.Println(er)
	// }

	
	_, er := config.BulkParser("config\\bulkconfig.json")

	if er != nil {
		fmt.Println(er)
	}


}

package config

import ("encoding/json"
"os"
"fmt")


func ParsHeader(pathJson string) (map[string]string, error) {
	
	file, err := os.ReadFile(pathJson)
	if err != nil{
		return nil, err
	}
	
	var header = make(map[string]string)
	err = json.Unmarshal(file, &header)
	if err != nil{
		return nil, err
	}

	for b,i:=range header{
		fmt.Printf("%v\t%v\n", b, i)
	}
	

	return header, nil 
}
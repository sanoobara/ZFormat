package config

import ("encoding/json"
"os"
"fmt")

type Config struct{

	InputFolder string 
	OutputFolder string
}

//Парсер заголовка 
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


func MainConfig(pathConfig string) (Config, error){
	var conf  Config
	file, err := os.ReadFile(pathConfig)
	if err != nil{
		return conf, err
	}
	
	
	err = json.Unmarshal(file, &conf)
	if err != nil{
		return conf, err
	}

	return conf, nil 

}



type Bulk struct{
	Zheader []string
	Service []string
}

func BulkParser(pathJson string) (Bulk, error){
	
	var Bstruct Bulk
	file, err := os.ReadFile(pathJson)
	if err != nil{
		return Bstruct, err
	}
	
	
	err = json.Unmarshal(file, &Bstruct)
	if err != nil{
		return Bstruct, err
	}

	for _, item := range Bstruct.Zheader{
		fmt.Println(item)
	}
	return Bstruct, nil
}

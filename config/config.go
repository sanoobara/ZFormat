package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"rm/config/cleanenv"
)
//Структура для парсинга основного конфигурационного файла
type Config struct {
	InputFolder  string
	OutputFolder string
}

//Локатион нужный полей в звголовке
type Location struct{
	startByte int
	stopByte int
}

// type head struct {
// 	Name      string
// 	Value     string
// 	Descr     string
// 	StartByte string
// 	StopByte  string
// }

// type Header struct {
// 	Items []head
// }
// 
func ParsLocation(pathJson string) (map[string]Location, error) {
	var header Header
	var maps = make(map[string]Location)

	file, err := os.ReadFile(pathJson)
	if err != nil {
		return maps, err
	}

	err = json.Unmarshal(file, &header)
	if err != nil {
		return maps, err
	}

	

	for _, item := range header.Items{
		start, err:=strconv.Atoi(item.StartByte)
		if err!=nil{
			fmt.Println(err)
		}
		stop, err:= strconv.Atoi(item.StopByte)
		if err!=nil{
			fmt.Println(err)
		}

		l:=  Location{start,stop }
		maps[item.Name] = l
	}
	
	

	return maps, nil
}

//Парсер заголовка


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

	return header, nil
}

// Парсер основого конфигурациооного файла
func MainConfig(pathConfig string) (Config, error) {
	var conf Config
	file, err := os.ReadFile(pathConfig)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil

}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type (

		ConfigParser struct {
			type Items []struct {
				Name      string
				Value     string
				Descr     string
				StartByte string
				StopByte  string

			}
		}

)
func parseJSON(r io.Reader, str interface{}) error {
	return json.NewDecoder(r).Decode(str)
	
func NewConfPars()(*ConfigParser, error) {
	cfg :=&ConfigParser{}

	err:= cleanenv.ReadConfig("./zheader.json", cfg)
	if err != nil{
		return nil, fmt.Errorf("config error: %w", err)
	}
	
	
	if err != nil {
		return nil, err
	}

	return cfg, nil	
	
}

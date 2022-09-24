package config

import (
	"encoding/json"
	"fmt"
	"os"
	
	"strconv"
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


//
type head struct {
	Name      string
	Value     string
	Descr     string
	StartByte string
	StopByte  string
}

type Header struct {
	Items []head
}

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
func ParsHeader(pathJson string) (Header, error) {
	var header Header
	file, err := os.ReadFile(pathJson)
	if err != nil {
		return header, err
	}

	err = json.Unmarshal(file, &header)
	if err != nil {
		return header, err
	}

	var maps = make(map[string]Location)

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
	fmt.Println(maps["OPR"].startByte)
	

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

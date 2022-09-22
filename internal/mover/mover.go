//Пакет для переноса файлов из одной директории в другую
//с предоставлением всей нужной информации о файле и заносе данных в базу данных

package mover

import (
	"fmt"
	"log"
	"os"
	"path"
)

func MoveFiles(pathInputFolder string, pathOutputFolder string) {

	DirEntry, err := os.ReadDir(pathInputFolder)
	if err != nil {
		log.Fatal(err)
	}

	var newPath string // Инициализация нового пути файла
	var oldPath string //
	for i := 0; i < len(DirEntry); i++ {

		oldPath = path.Join(pathInputFolder, DirEntry[i].Name())
		newPath = path.Join(pathOutputFolder, DirEntry[i].Name())
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File: %v has been rename %v", oldPath, newPath)
	}

}

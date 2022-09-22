//Пакет для переноса файлов из одной директории в другую
//с предоставлением всей нужной информации о файле и заносе данных в базу данных

package mover

import (
	"fmt"

	"os"
	"path"
)

func MoveFiles(pathInputFolder string, pathOutputFolder string) error {
	var err error
	err = nil
	DirEntry, err := os.ReadDir(pathInputFolder)
	if err != nil {
		return err
	}
	if len(DirEntry) == 0 {
		fmt.Println("Current folder is empty")
		return err
	}

	var newPath string // Инициализация нового пути файла
	var oldPath string //
	for i := 0; i < len(DirEntry); i++ {
		//Получим старый полный путь
		oldPath = path.Join(pathInputFolder, DirEntry[i].Name())
		// Проверим на наличие в директории директорий,
		//Если такие есть - удалим их
		if DirEntry[i].IsDir() == true {

			MoveFiles(oldPath, pathOutputFolder)
			err := os.Remove(oldPath)
			if err != nil {
				return err
			}
			fmt.Printf("Folder %v has been remove\n", oldPath)
		}
		//
		newPath = path.Join(pathOutputFolder, DirEntry[i].Name())
		if DirEntry[i].IsDir() == false {
			err = os.Rename(oldPath, newPath)
			if err != nil {

				return err
			}

			fmt.Printf("File: %v has been rename %v\n", oldPath, newPath)
		}

	}
	return err
}

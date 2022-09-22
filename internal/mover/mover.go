//Пакет для переноса файлов из одной директории в другую
//с предоставлением всей нужной информации о файле и заносе данных в базу данных

package mover

import (
	"fmt"

	"os"
	"path"
)

type OnesFile struct {
	FSold string //Корневая начальная директория
	FSnew string // Корневая конечная директория

	Name string
	Size int64
	ID   int
}

//Перенос файлов
func MoveFiles(files []OnesFile, outputFolder string) error {

	FS := outputFolder

	for _, f := range files {

		f.FSnew = FS

		newPath := path.Join(f.FSnew, f.Name)
		oldPath := path.Join(f.FSold, f.Name)

		err := os.Rename(oldPath, newPath)
		if err != nil {

			return err
		}

		fmt.Printf("File: %v has been rename %v\n", oldPath, newPath)
	}
	return nil

}

func GetListFiles(InputFolder string) ([]OnesFile, error) {

	//
	FS := InputFolder
	DirEntry, err := os.ReadDir(InputFolder)
	if err != nil {
		return nil, err
	}

	if len(DirEntry) == 0 {
		fmt.Println("Current folder is empty")
		return nil, nil
	}

	var Files []OnesFile
	for _, f := range DirEntry {

		//Извлечение иеформации о файле
		info, err := f.Info()
		if err != nil {
			return nil, err
		}
		// Проверка на директорию
		if info.IsDir() == false {
			var file OnesFile
			file.Name = info.Name()
			file.Size = info.Size()
			file.FSold = FS
			file.ID = os.Getuid()
			//Добавление в массив
			Files = append(Files, file)
		} else {
			continue
		}

	}
	return Files, nil
}

//Пакет для переноса файлов из одной директории в другую
//с предоставлением всей нужной информации о файле и заносе данных в базу данных

package mover

import (
	"fmt"
	"rm/internal/file"
	"os"
	"path"
)



//Перенос файлов
func MoveFiles(files []file.OnesFile, outputFolder string) error {

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

// Получение списка файлов в заданной директории
// Игнорирование директорий 
func GetListFiles(InputFolder string) ([]file.OnesFile, error) {

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

	var Files []file.OnesFile
	for _, f := range DirEntry {

		//Извлечение информации о файле
		info, err := f.Info()
		if err != nil {
			return nil, err
		}
		// Проверка на директорию
		if info.IsDir() == false {
			var file file.OnesFile
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

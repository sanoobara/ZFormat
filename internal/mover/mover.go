//Пакет для переноса файлов из одной директории в другую
//с предоставлением всей нужной информации о файле и заносе данных в базу данных

package mover

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

type OnesFile struct{
	
	FSold fs.FS //Корневая начальная директория
	FSnew fs.FS // Корневая конечная директория

	Name string
	Size string
	ID int

}

//Перенос файлов 
func MoveFiles([]OnesFile, pathOutputFolder string) error {

	var err error 
	err = nil
	var oldPath string //
	for file, _ := DirEntry {
		//Получим старый полный путь
		info := file.Info()
		oldPath = path.Join(os.DirFS(path), DirEntry[i].Name())
		// Проверим на наличие в директории директорий,
		//Если такие есть - удалим их
		if DirEntry[i].IsDir() == true {

			MoveFiles(oldPath, pathOutputFolder)
			err = os.Remove(oldPath)
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


func GetListFiles(pathInputFolder string) ([]OnesFile, error){
	

	FS := os.DirFS(pathInputFolder)
	DirEntry, err := os.ReadDir(pathInputFolder)
	if err != nil {
		return nil, err
	}


	if len(DirEntry) == 0 {
		fmt.Println("Current folder is empty")
		return nil, nil
	}

    var file OnesFile
	

	var Files = []OnesFile
	for f, _ := range DirEntry{
		
		if f.IsDir != true{
			file.Name = f.Name()
			file.Size = f.Size()
			file.FSold = FS
			
			
			Files = append(Files, file)
		}

	}
}

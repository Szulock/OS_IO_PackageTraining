package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Разработайте программу, которая выводит список всех файлов в директории.

func allFiles(directoryPath string) {
	directoryPath = filepath.FromSlash(directoryPath)

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println("Не получилось найти - неверный адрес")
	}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
		}
	}
}

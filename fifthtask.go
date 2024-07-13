package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Напишите программу для удаления всех файлов с определенным расширением в директории.

func delAllTxtFile(directoryPath string) {
	directoryPath = filepath.FromSlash(directoryPath)

	files, err := filepath.Glob(filepath.Join(directoryPath, "*"+"txt"))

	if err != nil {
		fmt.Println("Нет файлов с заданным расширением")
	}
	for _, file := range files {
		fmt.Println("Удаление файла:", file)
		err := os.Remove(file)
		if err != nil {
			fmt.Println("Ошибка при удалении файла:", err)
		}
	}

	fmt.Println("Все файлы с расширением TXT успешно удалены.")
}

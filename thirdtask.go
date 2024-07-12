package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Реализуйте функцию, проверяющую, существует ли файл или директория.

func IsExist(fileOrDirection string) bool {
	fileOrDirection = filepath.FromSlash(fileOrDirection)
	_, err := os.Stat(fileOrDirection)
	if os.IsNotExist(err) {
		fmt.Println("Такого не существует")
		return false
	} else {
		fmt.Println("Такой файл существует")
		return true
	}
}

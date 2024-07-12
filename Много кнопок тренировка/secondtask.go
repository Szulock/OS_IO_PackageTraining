package main

import (
	"fmt"
	"io"
	"os"
)

func Copyfile(OldFile, FreshFile string) {
	srcFile, err := os.Open(OldFile)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create("FreshFile")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Printf("Copied %d bytes from Example to FreshExample\n", bytesCopied)
}

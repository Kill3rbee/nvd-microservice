// create-directory
package main

import (
		"fmt"
		"os"
)

func createDirectory(dirName string) bool {
		src, err := os.Stat(dirName)

		if os.IsNotExist(err) {
				errDir := os.MkdirAll(dirName, os.ModePerm)
				if errDir != nil {
						panic(err)
				}
				return true
		}

		if src.Mode().IsRegular() {
				fmt.Println(dirName, "already exist as a file!")
				return false
		}

		return false
}

func main() {
	if len(os.Args) != 2 {
			fmt.Printf("Usage : %s <directory>\n", os.Args[0])
			os.Exit(0)
	}

	directory := os.Args[1]
	createDirectory(directory)
	// fmt.Println(directory, "created : ", result)
}
// download-files.go
package main

import (
    "bufio"
    "fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"os/exec"
	"path/filepath"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFromUrl(dirName string, url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// Call create-directory microservice
    if dirName != "" {
		path := filepath.Join(dirName, fileName)
		output, err := os.Create(path)
		
		if err != nil {
			fmt.Println("Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		_, err = io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}

    } else {
		output, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		_, err = io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
    }
}

// func createDirectory(dirName string) bool {
// 	src, err := os.Stat(dirName)

// 	if os.IsNotExist(err) {
// 			errDir := os.MkdirAll(dirName, os.ModePerm)
// 			if errDir != nil {
// 					panic(err)
// 			}
// 			return true
// 	}

// 	if src.Mode().IsRegular() {
// 			fmt.Println(dirName, "already exist as a file!")
// 			return false
// 	}

// 	return false
// }

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <directory>\n", os.Args[0])
		os.Exit(0)
	}

	directory := os.Args[1]

	// Call microservice for creating directory
	dirService := "../../bin/create-directory"
	args := []string{directory}

	result := exec.Command(dirService, args...)

	err := result.Start()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Command executed : ", result)

	// Call microservice for finding links on page
    linksService := "../../bin/find_links_in_page"
    cmd := exec.Command(linksService)

    stdout, _ := cmd.StdoutPipe()
    cmd.Start()

    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
		m := scanner.Text()
		url := "https://nvd.nist.gov/feeds/json/cve/1.1/" + m
		downloadFromUrl(directory, url)
    }
    cmd.Wait()
}
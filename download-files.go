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
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
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

	//fmt.Println(n, "bytes downloaded.")
}

func main() {
	// Call microservice for finding links on page
    args := "../bin/find_links_in_page"
    cmd := exec.Command(args)

    stderr, _ := cmd.StdoutPipe()
    cmd.Start()

    scanner := bufio.NewScanner(stderr)
    for scanner.Scan() {
		m := scanner.Text()
		url := "https://nvd.nist.gov/feeds/json/cve/1.1/" + m
		downloadFromUrl(url)
    }
    cmd.Wait()
}
// find_links_in_page.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "os"
)

func main() {
    if len(os.Args) != 3{
        fmt.Printf("Usage : %s <url> <regexp>\n", os.Args[0])
        os.Exit(0)
    }

    url := os.Args[1] //"https://nvd.nist.gov/vuln/data-feeds#JSON_FEED"
    useregexp := os.Args[2] //'nvdcve-1.1-[0-9]*\.json\.zip'
   
    // Make HTTP request
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Read response data in to memory
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal("Error reading HTTP body. ", err)
    }

    // Create a regular expression to find comments
    //nvdcve-1.1-[0-9]*\.json\.zip
    re := regexp.MustCompile(useregexp)
    comments := re.FindAllString(string(body), -1)
    if comments == nil {
        fmt.Println("No matches.")
    } else {
        for _, comment := range comments {
            fmt.Println(comment)
        }
    }
}
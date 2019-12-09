// find_links_in_page.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
)

func main() {
    // Make HTTP request
    response, err := http.Get("https://nvd.nist.gov/vuln/data-feeds#JSON_FEED")
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
    re := regexp.MustCompile(`nvdcve-1.1-[0-9]*\.json\.zip`)
    comments := re.FindAllString(string(body), -1)
    if comments == nil {
        fmt.Println("No matches.")
    } else {
        for _, comment := range comments {
            fmt.Println(comment)
        }
    }
}
package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"flag"
	"strings"
)

//make http request
// params: string, channel string
func GetRequest(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//converting the response body to string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	md5 := GenMD5(bodyString)
	//saving the output in channel
	ch <- url + " " + md5
}

//generating the md5 hash of http request
// params: string
// output: MD5 (hash) of the given string 
func GenMD5(response string) string {
	hasher := md5.New()
	hasher.Write([]byte(response))
	return hex.EncodeToString(hasher.Sum(nil))
}

//verifying the user provided flag value
func IsFlagExist(name string) bool {
    exist := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            exist = true
        }
    })
    return exist
}

//Normalising the url
// params: string
// output: string
func NormaliseUrl(url string) string{
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}

func main() {
	//defining the flag to indicate the limit of parallel request
	numbPtr := flag.Int("parallel", 10, "limit of parallel request")
	flag.Parse()
	//initializing the buffered channel
	ch := make(chan string, *numbPtr)
	//saving urls from user given arguments
	urls := os.Args[1:]
	if IsFlagExist("parallel") {
		fmt.Println(len(os.Args))
		if len(os.Args) > 2 {
			urls = os.Args[2:]
		} else {
			fmt.Println("please add arguments")
		}
    } 
	for _,url := range urls{
		url = NormaliseUrl(url)
		go GetRequest(url, ch)
	}
	for range urls{
		fmt.Println(<-ch) //printing the output and making the capcatity in channel
	}
}
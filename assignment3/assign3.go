package main

import (
	"fmt"
	"net/http"    
  	"io/ioutil"   
  	"io"
  	"time"
  	"os"
  	"bytes"
  	"encoding/json"
)


type Lang struct {
	name string
	urls string
	bytes uint
	time time.Duration
}

func Crawl(print func(Lang), lang Lang) {

	start := time.Now()
	
	resp, err := http.Get(lang.urls)  
	if err != nil {
		return
	}
                                               
  	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	count,err := io.Copy(ioutil.Discard, bytes.NewReader(body))
  	
  	if err != nil {
  		return
  	}

  	lang.bytes = uint(count)
  		
  	defer resp.Body.Close() 

	end := time.Now()

	lang.time = end.Sub(start)
	print(lang)

	return
}

func main() {
	printStruct := func (lang Lang) {
		fmt.Printf("%+v\n", lang)
	}

	printJSON := func (lang Lang) {
		data,err := json.Marshal(lang)
		if err == nil {
			os.Stdout.Write(data)
		}
	}
	python := Lang{name: "Python", urls: "https://www.python.org/"}
	Crawl(printStruct, python)
	Crawl(printJSON, python)

	ruby := Lang{name: "Ruby", urls: "https://www.ruby-lang.org/en/"}
	Crawl(printStruct, ruby)
	Crawl(printJSON, ruby)

	golang := Lang{name: "GoLang", urls: "https://golang.org/"}
	Crawl(printStruct, golang)
	Crawl(printJSON, golang)



}



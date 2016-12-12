package main

import (
	"fmt"
	"net/http"    
  	"io/ioutil"   
  	"io"
  	"time"
  	"bytes"
  	"sync"
)


type Lang struct {
	name string
	urls string
	bytes uint
	time time.Duration
}

func Crawl(lang Lang, wg sync.WaitGroup) {

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
	fmt.Printf("%+v\n", lang)
	defer wg.Done()

	return
}

func main() {
	

	var wg sync.WaitGroup

	python := Lang{name: "Python", urls: "https://www.python.org/"}
	ruby := Lang{name: "Ruby", urls: "https://www.ruby-lang.org/en/"}
	golang := Lang{name: "GoLang", urls: "https://golang.org/"}
	wg.Add(3);

	go Crawl(python,wg)
	go Crawl(ruby,wg)
	go Crawl(golang,wg)
	wg.Wait()

}



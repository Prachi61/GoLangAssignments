package main

import (
	"fmt"
	"net/http"    
  	"io/ioutil"   
  	"io"
  	"time"
  	"bytes"
  	//"sync"
)


type Lang struct {
	name string
	urls string
	bytes uint
	time time.Duration
}

type total struct {
	tbytes uint
	ttime time.Duration
}

func Crawl(lang Lang) Lang {
	//defer close(ch)
	//t := new(total)

	start := time.Now()
	
	resp, err := http.Get(lang.urls)  
	if err != nil {
		return lang
	}
                                               
  	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return lang
	}

	count,err := io.Copy(ioutil.Discard, bytes.NewReader(body))
  	
  	if err != nil {
  		return lang
  	}

  	lang.bytes = uint(count)
  		
  	defer resp.Body.Close() 

	end := time.Now()

	lang.time = end.Sub(start)
	/*for {
		t <- ch
		t.tbytes = t.tbytes + lang.bytes
		t.ttime = t.ttime + lang.time
		ch <- *t
	}
	//for res := range ch {
      //  t.tbytes = res.tbytes + lang.bytes
		//t.ttime = res.ttime + lang.time
    //}

    //ch <- *t
*/	
	fmt.Printf("%+v\n", lang)
	//wg.Done()

	return lang
}

func main() {
	

	ch := make(chan total,4)
	res := new(total)
	res.ttime = 0
	res.tbytes = 0
	ch <- *res
	python := Lang{name: "Python", urls: "https://www.python.org/"}
	ruby := Lang{name: "Ruby", urls: "https://www.ruby-lang.org/en/"}
	golang := Lang{name: "GoLang", urls: "https://golang.org/"}
	//wg.Add(3);

	lan1 := Crawl(python)
	res.ttime = res.ttime + lan1.time
	res.tbytes = res.tbytes + lan1.bytes
	fmt.Println(*res)

	lan2 := Crawl(ruby)
	res.ttime = res.ttime + lan2.time
	res.tbytes = res.tbytes + lan2.bytes
	fmt.Println(res)
	
	lan3 := Crawl(golang)
	res.ttime = res.ttime + lan3.time
	res.tbytes = res.tbytes + lan3.bytes
	fmt.Println(res)
	fmt.Println(<-ch)
	ch <- *res
	
	fmt.Println(<-ch)

}



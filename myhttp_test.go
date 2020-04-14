package main

import (
	"testing" 
	"time"
)

//Verifying the number of parallel request inside GetRequest function
func TestGetRequest(t *testing.T) {

	parallel := 5
	ch := make(chan string,parallel)
	urls := [...]string{"http://google.com", "http://fb.com", "http://gmail.com","http://hotmail.com", "http://quora.com","http://adjust.com", 
	"http://google.com", "http://fb.com", "http://gmail.com","http://hotmail.com", "http://quora.com","http://adjust.com", "http://sap.com"}
	for _,url := range urls{
		go GetRequest(url, ch)
	}

	time.Sleep(1 * time.Second)
	if parallel < len(ch){
		t.Errorf("Test failed: expected %d, got %d", parallel, len(ch))
	}
	
}
 
//Testing the normalise url function
 func TestNormaliseUrl(t *testing.T) {
	
	url := "http://google.com"
	ans := NormaliseUrl("google.com")
	if url != ans{
		t.Errorf("Test failed: expected %v, got %v", url, ans)
	}
}
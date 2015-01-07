package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var cb = NewCachedBytes([]byte("site loading"))
var chMaybeUpdate = make(chan chan struct{})

const targetHeight = 194300

func main() {
	go updater()

	// start http server
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8686", nil)
	if err != nil {
		log.Fatalf("error http.ListenAndServe: %v\n", err)
	}
}

func updater() {
	var diff int64
	for {
		// update
		chUpdateDone := make(chan struct{})
		chMaybeUpdate <- chUpdateDone
		log.Println("updating")
		func() {
			defer close(chUpdateDone)
			resp, err := http.Get("https://explorer.guldencoin.com/sapi/chainheight")
			if err != nil {
				log.Printf("error getting chain height: %v\n", err)
				return
			}
			defer resp.Body.Close()
			heightBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("error getting chain height: %v\n", err)
				return
			}
			height, _ := strconv.ParseInt(string(heightBytes), 10, 64)
			diff = targetHeight - height
			var msg string
			if diff < 0 {
				//++ link to explorer
				msg = fmt.Sprint("Digi is released!\n")
			} else {
				msg = fmt.Sprintf("Digi is coming in %d blocks. Estimated time: %s\n", diff, time.Now().Add(time.Duration(diff*150)*time.Second).String())
			}
			cb.Update([]byte(msg))
		}()

		// sleep a while before new update is possible
		// sleep duration depends on how far away the change is
		if diff > 48 {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	select {
	case chUpdateDone := <-chMaybeUpdate:
		<-chUpdateDone
	default:
	}
	cb.RLock()
	cb.WriteTo(w)
	cb.RUnlock()
	log.Printf("served %s\n", r.RemoteAddr)
}

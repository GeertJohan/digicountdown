package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var cb = NewCachedBytes([]byte("site loading"))
var chMaybeUpdate = make(chan chan struct{})
var tmplPage *template.Template

func init() {
	var err error
	tmplPage, err = template.New("digicountdown").Parse(page)
	if err != nil {
		log.Fatalf("error parsing page template: %v", err)
	}
}

const targetHeight = 194300

func main() {
	go updater()

	// start http server
	http.HandleFunc("/style.css", styleHandler)
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
			newDiff := targetHeight - height
			if newDiff == diff {
				return
			}
			diff = newDiff
			data := &pageData{}
			if diff < 0 {
				data.Activated = true
			} else {
				duration := time.Duration(diff*150) * time.Second
				data.Blocks = diff
				data.Duration = fmt.Sprintf("%.0fd %dh %dm", duration.Hours()/24, int(duration.Hours())%24, int(duration.Minutes())%60) // int(duration.Seconds())%60
				data.EstimatedTime = time.Now().Add(duration).Format("Mon 2006-01-02 15:04:05 -0700 MST")                               // Mon Jan 2 15:04:05 -0700 MST 2006
			}
			pageBuffer := &bytes.Buffer{}
			err = tmplPage.Execute(pageBuffer, data)
			if err != nil {
				log.Printf("error executing template: %v", err)
				return
			}
			cb.Update(pageBuffer.Bytes())
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

func styleHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	w.Header().Set("Content-Type", "text/css")
	w.Write(style)
}

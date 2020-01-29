package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ffprobe "github.com/vansante/go-ffprobe"
)

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Ok"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.WriteHeader(400)
		w.Write([]byte("Url Missing"))
	}
	data, err := ffprobe.GetProbeData(url, 2000*time.Millisecond)
	if err != nil {
		log.Printf("Error getting data: %v", err)
		w.WriteHeader(400)
		w.Write([]byte("Some Error Occurred"))
	}

	duration := fmt.Sprintf("%f", data.Format.Duration().Seconds())
	w.WriteHeader(200)
	w.Write([]byte("{\"duration\":" + duration + "}"))
}

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/duration", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

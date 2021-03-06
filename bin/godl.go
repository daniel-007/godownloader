package main

import (
	"crawler/downloader"
	"flag"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	port := flag.String("port", "8100", "port number")
	mode := flag.String("mode", "download", "mode")
	flag.Parse()

	downloader.Port = *port
	//downloader.LogWriter, _ = os.Create("access.log")

	if *mode == "download" {
		http.Handle("/download", downloader.NewDownloadHanler())
	}
	if *mode == "simple" {
		http.Handle("/simple", downloader.NewSimpleDownloadHandler())
	}
	if *mode == "image" {
		http.Handle("/image", downloader.NewPicDownloadHandler())
	}
	if *mode == "realtime" {
		http.Handle("/realtime", downloader.NewRealtimeDownloadHandler())
	}
	if *mode == "redirect" {
		http.Handle("/redirect", downloader.NewRedirectorHandler())
	}
	if *mode == "filter" {
		http.Handle("/filter", downloader.NewBloomFilterHandler())
	}

	s := &http.Server{
		Addr:           ":" + *port,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

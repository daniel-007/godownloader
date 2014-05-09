package downloader

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type RealtimeDownloadHandler struct {
	Downloader                     *HTTPGetDownloader
}

func (self *RealtimeDownloadHandler) ProcessLink(link string) string {
	log.Println(time.Now().Unix(), "downloader", "start", link)
	html := ""
	resp := ""
	var err error	

	html, resp, err = self.Downloader.Download(link)
	if err != nil {
		return ""
	}

	if len(html) < 100 {
		return ""
	}

	if !IsChinesePage(html) {
		return ""
	}
	log.Println(time.Now().Unix(), "downloader", "finish", link)
	page := WebPage{Link: link, Html: html, RespInfo: resp, DownloadedAt: time.Now().Unix()}

	ret := strconv.FormatInt(page.DownloadedAt, 10) + "\t" + page.Link + "\t" + page.Html + "\t" + page.RespInfo
	return ret
}


func NewRealtimeDownloadHandler() *RealtimeDownloadHandler {
	ret := RealtimeDownloadHandler{}
	ret.Downloader = NewHTTPGetDownloader()
	return &ret
}

func (self *RealtimeDownloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	link := req.FormValue("link")

	fmt.Fprint(w, self.ProcessLink(link))
}
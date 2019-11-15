package v1

import (
	"bytes"
	"fmt"
	"github.com/anielski/download-url-in-go/app"
	"github.com/anielski/download-url-in-go/handlers"
	"github.com/anielski/download-url-in-go/models"
	"io"
	"net/http"
	"time"
)

type FetcherWork struct {
	Id       uint
	Url      string
	Interval uint
	Delete   bool
}

//is like worker's heap
var FetchersWork []*FetcherWork

//Start workers when app run
//@todo change this function as Singleton and secure against restarting @see sync.Once
func StartWorker(g handlers.Gwp) {
	app.Logger(nil).Info("START")
	fitchers, err := models.GetFetchers(app.API.DB)
	if err != nil {
		fmt.Println(err)
	}
	app.Logger(nil).Info("URL: ", fitchers)

	for _, fitcher := range fitchers {
		AddFetcher(fitcher.Id, fitcher.Url, fitcher.Interval)
	}

}

//Edit fetcher as work
func EditFetcher(id uint, url string, interval uint) {
	app.Logger(nil).Info("URL: ", url)
	//fmt.Println(url)
	for _, fetcher := range FetchersWork {
		if (fetcher.Id == id) {
			fetcher.Url = url
			fetcher.Interval = interval
		}
	}
}

//Add fetcher as work
func AddFetcher(id uint, url string, interval uint) {
	app.Logger(nil).Info("URL: ", url)
	fmt.Println(url)
	fw := &FetcherWork{
		Id:       id,
		Url:      url,
		Interval: interval,
		Delete:   false,
		//Channel: make(chan bool),
	}
	FetchersWork = append(FetchersWork, fw)
	go job(fw)
}

//Mark job as delete
func DeleteFetcherWork(id uint) {
	for _, fetcher := range FetchersWork {
		if (fetcher.Id == id) {
			fetcher.Delete = true
		}
	}
}

func deleteFetcherWork(id uint) {
	for key, fetcher := range FetchersWork {
		if (fetcher.Id == id) {
			if (key != len(FetchersWork)-1) {
				FetchersWork[key] = FetchersWork[len(FetchersWork)-1]
			}
			FetchersWork[len(FetchersWork)-1] = nil
			FetchersWork = FetchersWork[:len(FetchersWork)-1]
			break
		}
	}
}

func job(fw *FetcherWork) {
	for {
		if fw.Delete {
			deleteFetcherWork(fw.Id)
			break
		}
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		start := time.Now()
		resp, err := client.Get(fw.Url)
		if err != nil {
			fmt.Println(err)
		} else {
			var bodyBytes bytes.Buffer
			//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
			//bodyBytes, err := ioutil.ReadAll(resp.Body)
			if _, err := io.Copy(&bodyBytes, resp.Body); err != nil {
				fmt.Printf("za duża wartość %s: %v \n", fw.Url, err)
			}
			nbytes := bodyBytes.Len()
			if err != nil {
				fmt.Printf("podczas odczytywania %s: %v \n", fw.Url, err)
			} else if (nbytes > 1048576) {
				fmt.Printf("za duża wartość %s: %v \n", fw.Url, err)
			} else if resp.StatusCode == http.StatusOK {
				secs := time.Since(start).Seconds()
				fmt.Printf("%.2fs %7d  %s \n", secs, nbytes, fw.Url)
				models.SaveHistory(fw.Id, bodyBytes.Bytes(), secs, app.API.DB)
			} else {
				fmt.Printf("Error \n")
			}
			resp.Body.Close()
		}
		time.Sleep(time.Duration(fw.Interval) * time.Second)
	}
}

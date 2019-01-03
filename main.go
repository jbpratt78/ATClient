package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type streamList struct {
	Strims       []stream `json:"stream_list"`
	Streams      int      `json:"streams"`
	TotalViewers int      `json:"total_viewers"`
}

type stream struct {
	Username string `json:"username"`
	Viewers  int    `json:"viewers"`
}

func main() {
	url := "https://angelthump.com/api"

	angelthumpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "chibidesti")

	res, getErr := angelthumpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	// careful, ReadAll() can hang waiting on response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	streams := streamList{}
	jsonErr := json.Unmarshal(body, &streams)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println("Total viewer count: ", streams.TotalViewers)
	fmt.Println("Stream count: ", streams.Streams)
	fmt.Println(streams.Strims)

}

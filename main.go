package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://distopia.savi2w.workers.dev/go"

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Algo deu errado:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Algo deu errado: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusFound {
		distopiaHeader := resp.Header.Get("Distopia")
		fmt.Println("Distopia header value:", distopiaHeader, resp.StatusCode)
	} else {
		fmt.Println("Status Code:", resp.StatusCode, "is not a redirect.")
	}
}

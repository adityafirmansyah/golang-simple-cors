package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/cors"
)

func getRemoteHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if htmlStr, err := getRemoteHTML(url); err != nil {
			fmt.Fprintf(w, "Failed to get page: %v", err)
		} else {
			fmt.Fprintf(w, htmlStr)
		}
	})

	handler := cors.Default().Handler(mux)
	panic(http.ListenAndServe(":9000", handler))
}

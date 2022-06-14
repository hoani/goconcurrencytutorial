package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func Ping(c *http.Client, url string) error {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return err
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	if res == nil || res.StatusCode != 200 {
		return errors.New("bad ping")
	}
	return nil
}

func main() {
	servers := []string{
		"http://hoani.net",
		"http://dronedeploy.com",
		"http://google.com",
		"http://ubuntu.com",
		"http://apple.com",
		"http://rickastley.co.uk",
	}

	start := time.Now()

	results := map[string]error{}

	c := &http.Client{}
	// Ping all servers.
	for _, server := range servers {
		result := Ping(c, server)
		results[server] = result
	}

	// Print results.
	for server, err := range results {
		fmt.Printf("server: %s\t", server)
		if err != nil {
			fmt.Printf("got error %v\n", err)
		} else {
			fmt.Println("success")
		}
	}
	fmt.Printf("\n Took %v \n", time.Since(start))
	fmt.Printf("\n Running go routines %v \n", runtime.NumGoroutine())
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func PingServer(c *http.Client, url string) error {
	res, err := c.Head(url)
	if err != nil {
		return err
	}
	if res == nil || res.StatusCode != 200 {
		return errors.New("call failed")
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
		result := PingServer(c, server)
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
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("\n Running go routines %v \n", runtime.NumGoroutine())
}

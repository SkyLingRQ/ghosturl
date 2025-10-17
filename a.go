package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func singleUrl(client *http.Client, url string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() { <-sem }()

	waybackConsult := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=*.%s/*&output=text&fl=original&collapse=urlkey", url)

	wayback, err := http.NewRequest("GET", waybackConsult, nil)
	if err != nil {
		fmt.Println("[ CREATE REQUEST ] ¡Ups! A error ocurred.")
		return
	}

	response, err := client.Do(wayback)
	if err != nil {
		fmt.Println("[ RESPONSE ] ¡Ups! A error ocurred.")
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[ READ HTML ] ¡Ups! A error ocurred.")
		return
	}

	fmt.Println(string(result))

}

func multipleUrls(url string, wg *sync.WaitGroup, sem chan struct{}) {

}

func main() {
	var wg sync.WaitGroup
	var client = &http.Client{
		Timeout: 3 * time.Minute,
	}
	url := flag.String("d", "", "Domain for the scan.")
	file := flag.String("f", "", "File for the scan.")
	flag.Parse()

	sem := make(chan struct{}, 100)

	if (*url == "" && *file == "") || (*url != "" && *file != "") {
		flag.Usage()
		os.Exit(1)
	}

	if *url != "" {
		wg.Add(1)
		go singleUrl(client, *url, &wg, sem)
	}

	if *file != "" {
		content, err := os.Open(*file)
		if err != nil {
			fmt.Println("[ READ FILE ] ¡Ups! A error ocurred.")
			return
		}

		defer content.Close()

		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			url := scanner.Text()
			wg.Add(1)
			go singleUrl(client, url, &wg, sem)
		}

	}

	wg.Wait()

}

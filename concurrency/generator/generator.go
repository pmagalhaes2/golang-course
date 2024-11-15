package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// <-chan - canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)</title>`)
			match := r.FindStringSubmatch(string(html))

			if len(match) > 1 {
				c <- match[1]
			} else {
				c <- fmt.Sprintf("No title found for %s", url)
			}
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://google.com")
	t2 := titulo("https://www.amazon.com", "https://youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}

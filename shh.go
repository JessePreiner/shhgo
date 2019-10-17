package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func parse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse(c)
	}
}

func crawl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("bad url: ", url)
		return
	}

	b := resp.Body

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("can't crawl: ", url)
		return
	}

	html, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Println("can't crawl: ", url)
		return
	}

	parse(html)

	defer b.Close()

}

func main() {
	url := `http://www.shush.se/index.php?shows`
	crawl(url)
}


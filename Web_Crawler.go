package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type FetchedSite struct {
	sites map[string]int
	mux sync.Mutex
}

func (s *FetchedSite) Regsiter(key string) {
	s.mux.Lock()
	s.sites[key]++
	s.mux.Unlock()
}

func (s *FetchedSite) isRegisterd(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.sites[key]
}

func Crawl(url string, depth int, fetcher Fetcher) {
 	sites := FetchedSite{sites: make(map[string]int)}

	var crawler = func(string, int, Fetcher, chan int){}
	crawler = func(url string, depth int, fetcher Fetcher, ch chan int) {
		fmt.Println("Enter crawler func: ", url)
		defer close(ch)
		if depth <= 0 {
			fmt.Println("exit crawler func (depth<=0): ", url)
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			fmt.Println("exit crawler func (err): ", url)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		chs := make([]chan int, len(urls))
		for i, u := range urls {
			if sites.isRegisterd(u) == 0 {
				sites.Regsiter(u)
				chs[i] = make(chan int)
				fmt.Println("open channel: ", chs[i], u)
				go crawler(u, depth-1, fetcher, chs[i])
			}
		}
		//waiting for all channels completion.
		for i := range chs {
			if chs[i] != nil {
				<-chs[i]
				fmt.Println("channel is closed: ", chs[i])
			}
		}

		fmt.Println("Exit crawler func: ", url)
		return
	}

	ch := make(chan int)
	sites.Regsiter(url)
	crawler(url, depth, fetcher, ch)
	<-ch
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

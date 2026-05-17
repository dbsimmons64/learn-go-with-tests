package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// A channel is a "typed pipe" - One goroutine sends values into a channel,
	// another receives them.
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	// blah
	return results
}

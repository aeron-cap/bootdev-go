package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(w WebsiteChecker, urls []string ) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, w(url)}
		}()

		for i := 0; i < len(resultChannel); i++ {
			r := <-resultChannel
			results[r.string] = r.bool
		}
	}
	return results
}

/*
 * we usually solve race conditions using channels
 * which can be instantiated by make(chan TYPE)
 * and <- is a send statement which is used to send
 * values like this make(chan TYPE) <- result TYPE
 * called a send statement
 *
 * for received statement we do
 * x := <- chanVariable
 * and then we can access x as values
 */
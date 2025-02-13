package fuck

import (
	"context"
	"log"
	"github.com/chromedp/chromedp"
	"strings"
	"golang.org/x/net/html"
)

// FuzzPayloads contains test payloads for XSS sanitization testing.



/*
func Fuzz(data []byte) int {
	png.Decode(bytes.NewReader(data))
	return 0
}
*/


func Fuzz(data []byte) int {
	// fuzz(ctx, "http://example.com/test", p)
	ctx, _ := chromedp.NewContext(context.Background())
	// defer cancel()
	fuzz_prog(ctx, "http://example.com/test", string(data));
	//png.Decode(bytes.NewReader(data))
	return 0
}


/*
func main() {
	// Target URL to fuzz
	targetURL := "http://example.com/test"

	// Create a context for Chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a wait group for parallel fuzzing
	var wg sync.WaitGroup

	// Start fuzzing with all payloads
	for _, payload := range FuzzPayloads {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			fuzz(ctx, targetURL, p)
		}(payload)
	}

	// Wait for all fuzzing routines to complete
	wg.Wait()

	fmt.Println("Fuzzing complete.")
}
*/












func is_allowed_html(text string) bool {

	tkn := html.NewTokenizer(strings.NewReader(text))

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return true
		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "h1" {
				continue
			} else {
				return false
			}
		}
	}
}









func fuzz_prog(ctx context.Context, url, payload string) {
	// Create a new context with timeout for each payload
	// localCtx, cancel := context.WithTimeout(ctx, 10000000000000)
	// defer cancel()

	localCtx, _ := context.WithTimeout(ctx, 10000000000000)
	// defer cancel()

	var result string
	
	if (!(is_allowed_html(payload))) {
		// Not allowed, therefore just return here already.
		return
	}

	// Define the steps for Chromedp to execute

	// panic("Fuck!");

	err := chromedp.Run(localCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#input`, chromedp.ByID),
		chromedp.SendKeys(`#input`, payload, chromedp.ByID),
		chromedp.Click(`#submit`, chromedp.ByID),
		chromedp.Text(`#output`, &result, chromedp.ByID),
	)

	if err != nil {
		log.Printf("Error running payload '%s': %v", payload, err)
		return
	}

	// Analyze the result
	if detectXSS(result) {
		//log.Printf("Potential XSS vulnerability detected with payload: %s", payload)
		//writeToFile(payload, result)
		panic("Found XSS vector");
	}
}

func detectXSS(response string) bool {
	// Basic check for unsanitized script execution or alerts
	return containsAny(response, []string{"<script>", "alert", "<img", "<svg", "<iframe"})
}

func containsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
		if contains(s, substr) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
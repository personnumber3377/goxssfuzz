package fuck

import (
	"log"
	"strings"
	"golang.org/x/net/html"
	"io/ioutil"
	//"os/exec"
	"bufio"
	"os"
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
	//ctx, _ := chromedp.NewContext(context.Background())
	// defer cancel()
	fuzz_prog(string(data));
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


func checkCharacterInFile(filePath string, targetChar byte) (bool, error) {
	// Open file and scan line-by-line
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if contains(scanner.Text(), targetChar) {
			return true, nil
		}
	}
	return false, scanner.Err()
}

// Helper function to check if targetChar is in the line
func contains(line string, targetChar byte) bool {
	for i := 0; i < len(line); i++ {
		if line[i] == targetChar {
			return true
		}
	}
	return false
}






func fuzz_prog(payload string) {
	// Create a new context with timeout for each payload
	// localCtx, cancel := context.WithTimeout(ctx, 10000000000000)
	// defer cancel()




	//localCtx, _ := context.WithTimeout(ctx, 10000000000000)
	// defer cancel()

	//var result string
	
	if (!(strings.Contains(payload, "alert(1)"))) {
		return
	}

	if (!(is_allowed_html(payload))) {
		// Not allowed, therefore just return here already.
		return
	}

	// Define the steps for Chromedp to execute

	// panic("Fuck!");
	/*
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
	*/




	// Analyze the result
	/*
	if detectXSS(result) {
		//log.Printf("Potential XSS vulnerability detected with payload: %s", payload)
		//writeToFile(payload, result)
		panic("Found XSS vector");
	}
	*/
	


	foundf, err := checkCharacterInFile("result.txt", byte('1'));

	if (foundf) {
		panic("oooffff");
	}

	erree := ioutil.WriteFile("input.txt", []byte(payload), 0644)
	if erree != nil {
		log.Fatal(err)
	}

	// cmd := exec.Command("python3", "xss.py") // You can change the command and its arguments accordingly

	// Run the command and capture the output
	// cmd.CombinedOutput()


	// Ok, so now read the result and then see the thing..

	found, _ := checkCharacterInFile("result.txt", byte('1'));

	if (found) {
		panic("oooffff");
	}


}


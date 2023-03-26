/*
#golang #gem: regexp can use functions for substitution.

re := regexp.MustCompile(`\$[A-Z_]+`)
urlTemplate := `https://$HOST:$PORT`

	conf := map[string]string{
		"HOST": "www.353solutions.com",
		"PORT": "443",
	}

	sub := func(match string) string {
		return conf[match[1:]]
	}

url := re.ReplaceAllStringFunc(urlTemplate, sub)
fmt.Println(url)
// https://www.353solutions.com:443
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\$[A-Z_]+`)
	urlTemplate := `https://$HOST:$PORT`

	conf := map[string]string{
		"HOST": "www.353solutions.com",
		"PORT": "443",
	}

	sub := func(match string) string {
		return conf[match[1:]]
	}

	url := re.ReplaceAllStringFunc(urlTemplate, sub)
	fmt.Println(url)
	// https://www.353solutions.com:443
}

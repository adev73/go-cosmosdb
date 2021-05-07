package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adev73/go-cosmosdb/cosmosapi"
)

// Format a JavaScript-file for inline use in a JSON file.
// Usage: cat some-script.js | [this cmd] ( | clipboard)
func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	sourceCode := cosmosapi.EscapeJavaScript(bytes)

	fmt.Fprintf(os.Stdout, sourceCode)
}

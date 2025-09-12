package prettier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"tvguide/textops"
)

// Unfortuantely no Named parameter for toFile := false
func PrintFromObject(content any, toFile bool) {
	// This is used to print object in json form, NOT for printing binary objects from response body in GET
	prettyJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		panic(err)
	}
	if toFile {
		textops.WriteFile(string(prettyJSON), "output.txt")
	} else {
		fmt.Println(string(prettyJSON))
	}
}

func PrintFromResponseBody(content []byte) {
	// This is for printing binary objects from response body in GET
	// Convert string to bytes
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, content, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(prettyJSON.String())
}

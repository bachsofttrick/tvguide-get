package prettier

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintFromObject(content any) {
	// This is used to print object in json form, NOT for printing binary objects from response body in GET
	prettyJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyJSON))
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

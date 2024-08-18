package prettier

import (
	"encoding/json"
	"fmt"
)

func Print(content any) {
	prettyJSON, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyJSON))
}

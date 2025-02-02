package utils

import (
	"encoding/json"
	"fmt"
)

func PrintToJSON(val any) {
	data, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(data))
}

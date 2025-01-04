package utils

import (
	"encoding/json"
	"fmt"
)

func PrintToJSON(val any) {
	data, _ := json.Marshal(val)
	fmt.Println(string(data))
}

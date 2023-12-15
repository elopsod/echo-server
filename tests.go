package main

import (
	"fmt"
	"github.com/caitlinelfring/go-env-default"
)

//	func PrettyString(str string) (string, error) {
//		var prettyJSON bytes.Buffer
//		if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
//			return "", err
//		}
//		return prettyJSON.String(), nil
//	}
func main() {
	//md := "{'k': 11}"
	//jsonString, _ := json.Marshal(md)
	//prettyJson, _ := PrettyString(string(jsonString))
	//fmt.Println(prettyJson)
	httpPort := env.GetDefault("HTTPs_PORT", "8080")
	fmt.Println(httpPort)
}

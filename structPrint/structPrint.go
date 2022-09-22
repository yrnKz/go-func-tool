package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

}

func StructPrint(s interface{}) {
	bs, _ := json.Marshal(s)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("%v", out)
}

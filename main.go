package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// mydata is json strusture
type Mydata struct {
	Name string
	Mail string
	Tel  string
}

// Str get string value
func (m *Mydata) str() string {
	return "<\"" + "\" " + m.Mail + ", " + m.Tel + ">"
}

func main() {
	// accese URL
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	
	//isGet information and error handring
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}

	defer re.Body.Close()

	// Extracting all contents
	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	// 　change to value of Golang
	var data []interface{}
	er = json.Unmarshal(s, &data)
	if er != nil {
		panic(er)
	}

	// Extracting data from an array with a for statement
	for i, im := range data {
		// in type assertion
		// example: org := value.(type)
		m := im.(map[string]interface{})
		fmt.Println(i, m["name"].(string), m["mail"].(string),m["tel"].(string))
	}
}
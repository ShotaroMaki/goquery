package main

import "github.com/PuerkitoBio/goquery"

func main() {
	//connnect url
	P := "https://golang.org"
	
	//target
	doc, er := goquery.NewDocument(P)
	
	// error handling
	if er != nil {
		panic(er)
	}
	// fetching <a> element
	doc.Find("a").Each(func(n int, sel *goquery.Selection){
		lk, _ := sel.Attr("href")
		
		// output to text
		println(n, sel.Text(), "(", lk, ")")
	})
}
package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	var pagename string
	pagename = "TestPage 1"
	p1 := &Page{Title: pagename, Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage(pagename)
	fmt.Println(string(p2.Body))
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

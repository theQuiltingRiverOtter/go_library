package main

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	pages  int
	read   bool
	owner  *Person
}

type Person struct {
	name string
	age  int
}

func (book Book) String() string {
	if book.owner != nil {
		return fmt.Sprintf("{Title: %+v Author: %+v Pages: %+v Read: %+v Owner: %+v}", book.title, book.author, book.pages, book.read, *book.owner)
	}
	return fmt.Sprintf("Title: %v, Author: %v, Pages: %v, Read: %v, Owner: nil", book.title, book.author, book.pages, book.read)

}

func (b *Book) UpdateTitle(s string) {
	b.title = s
}

func (b *Book) UpdateAuthor(s string) {
	b.author = s
}

func (b *Book) UpdatePages(n int) {
	b.pages = n
}

func (b *Book) UpdateRead() {
	b.read = !b.read
}

func (b *Book) UpdateOwner(o *Person) {
	if o != nil {
		b.owner = o
	}
}

func main() {
	o := &Person{"Megan", 39}
	c := Book{"Cat in the Hat", "Dr.Seuss", 20, false, o}

	fmt.Printf("%+v\n", c)

}

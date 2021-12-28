package main
import "fmt"

type Books struct{
	title string
	id int
}

func main(){
	var book1 Books
	var book2 Books
	book1.title="Wings of Fire"
	book1.id=123445677

	book2.title="Beyond the blue mountain"
	book2.id=375548627468


	// fmt.Printf("book1 title : %s\n",book1.title)
	// fmt.Printf("book1 id %d\n",book1.id)

	// fmt.Printf("book2 title :  %s\n",book2.title)
	// fmt.Printf("book2 id %d\n",book2.id)

	printBook(book1)
	printBook(book2)


}

func printBook(Book Books){
	fmt.Printf("Book title : %s\n",Book.title)
	fmt.Printf("Book id : %d\n",Book.id)
}
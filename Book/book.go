package book


//Book is an Entity/Business Object in Library Domain
type Book struct {
	BookID int
	BookTitle string
	BookAuthor string
	BookSubject string
}

func(b *Book) Title() string {
	return b.BookTitle
}


func(b *Book) UpdateTitle(newTitle string) bool {
	b.BookTitle = newTitle
	return true
}
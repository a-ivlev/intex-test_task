package storage

type BookModel struct {
	Title  string
	Author string
	Cost   int
}


type DB interface {
	GetBooksByAuthor(author string, result []*BookModel) ([]*BookModel, error)
}


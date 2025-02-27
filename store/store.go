package store

type Book struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Author []string `json:"author"`
	Press  string   `json:"press"`
}

type Store interface {
	Create(book *Book) error
	Update(book *Book) error
	Get(Id string) (Book, error)
	GetAll() ([]Book, error)
	Delete(Id string) error
}

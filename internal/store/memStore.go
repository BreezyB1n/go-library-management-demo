package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"fmt"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*store.Book
}

func (m *MemStore) Create(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, exist := m.books[book.Id]; exist {
		return fmt.Errorf("book already exist")
	}

	m.books[book.Id] = book
	return nil
}

func (m *MemStore) Update(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, exist := m.books[book.Id]; !exist {
		return fmt.Errorf("book not exist")
	}
	m.books[book.Id] = book
	return nil
}

func (m *MemStore) Get(Id string) (store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	book, exists := m.books[Id]
	if !exists {
		return store.Book{}, fmt.Errorf("book with ID %s does not exist", Id)
	}

	return *book, nil
}

func (m *MemStore) GetAll() ([]store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	var getBooks []store.Book
	for _, book := range m.books {
		getBooks = append(getBooks, *book)
	}
	return getBooks, nil
}

func (m *MemStore) Delete(Id string) error {
	m.Lock()
	defer m.Unlock()

	if _, exist := m.books[Id]; !exist {
		return fmt.Errorf("book not exist")
	}

	delete(m.books, Id)
	return nil
}

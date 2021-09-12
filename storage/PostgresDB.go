package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"strconv"
	"sync"
	"test-task-intech/config"
)

var ErrDBunavailable = errors.New("Базаданных не доступна")

type Service struct {
	mutex *sync.RWMutex
	Pool   []*pgx.Conn
	IsInit bool
}


func InitService(config *config.Config) DB {
	s := new(Service)
	s.mutex = &sync.RWMutex{}
	var backgroundTask = func() {
		var databaseUrl = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			config.DB.Username,
			config.DB.Password,
			config.DB.Host,
			config.DB.Port,
			config.DB.Database)

		for i := 1; i <= 10; i++ {
			conn, err := pgx.Connect(context.Background(), databaseUrl)
			if err != nil {
				log.Fatal("Ошибка при подключении к базе по URL = " + databaseUrl, err)
			}
			s.Pool = append(s.Pool, conn)
		}
	}
	go backgroundTask()
	s.IsInit = true
	return s
}

func (s Service) GetBooksByAuthor(author string, result []*BookModel) ([]*BookModel, error) {
	if s.IsInit != true {
		return nil, ErrDBunavailable
	}
	var conn *pgx.Conn
	for _, x := range s.Pool {
		if !x.IsClosed() {
			conn = x
			break
		}
	}
	s.mutex.RLock()
	rows, err := conn.Query(context.Background(), "select title, cost from books where author=$1", author)
	if err != nil {
		return nil, err
		log.Println("Не удалось получить книги по автору", err)
	}
	s.mutex.RUnlock()
	for rows.Next() {
		var title string
		var cost int
		err = rows.Scan(&title, &cost)
		if err == nil {
			result = append(result, &BookModel{title, author, cost})
		}
	}
	log.Println("Успешно выполнен запрос, заполнено записей: " + strconv.Itoa(len(result)))
	return result, nil
}
package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-task-intech/storage"
)

func GetByAuthorHandler(db storage.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author := vars["author"]
	var result = make([]*storage.BookModel, 0, 10)
	// Нет никакого смысла передавать author по указателю. По указателю лучше передавать слайс.
	// Так он будет занимать константное место в памяти, и не нужно думать, что скопируется, а что нет.
	result, err := db.GetBooksByAuthor(author, result)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
}

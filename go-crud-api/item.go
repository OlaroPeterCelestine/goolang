package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func getItems(w http.ResponseWriter, r *http.Request) {
	rows, _ := DB.Query("SELECT id, title FROM items")
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		rows.Scan(&i.ID, &i.Title)
		items = append(items, i)
	}
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	res, _ := DB.Exec("INSERT INTO items (title) VALUES (?)", item.Title)
	id, _ := res.LastInsertId()
	item.ID = int(id)

	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	DB.Exec("UPDATE items SET title = ? WHERE id = ?", item.Title, id)
	item.ID, _ = strconv.Atoi(id)
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	DB.Exec("DELETE FROM items WHERE id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}
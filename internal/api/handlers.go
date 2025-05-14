package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"example.com/internal/models"
	"example.com/internal/repository"

	"github.com/go-chi/chi/v5"
)

//create handlers For Specifioc Endpoints

//create a handler for the TODO endpoints / to Access TodoRepo Directly

type TodoHandler struct {
	todoRepo *repository.MongoTodoRepository
}

//Create A TodoHandler // basically its a Fucntion Returning a TodoHandler

func NewTodoHandler(todoRepo *repository.MongoTodoRepository) *TodoHandler {
	return &TodoHandler{todoRepo: todoRepo}
}

//create a handler for the /api/v1/todos endpoint

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	//get Todo FROM body
	var todo *models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error decoding todo:", err)
		return
	}
	//add to DB
	if err := h.todoRepo.CreateTodo(context.Background(), todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//return 201 created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//get id from url
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	//Delete A Todo Fro[m Db]
	if err := h.todoRepo.DeleteTodo(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//return 200 ok

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo Deleted Successfully"))
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	//get id from url
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	//Look up into Db
	todo, err := h.todoRepo.GetTodo(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//return 200 ok
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo) //Directly Convert Todo to Json
	/*
		other way to Do the Same is
		json,err:=json.Marshal(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)

	*/

}

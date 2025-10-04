package handler

import (
	"net/http"

	"github.com/luckmade/shorter-url/models"
)

type Handler struct {
	users models.UsersService
}

func NewHandler(us models.UsersService) *Handler {
	return &Handler{
		users: us,
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

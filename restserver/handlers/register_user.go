package handlers

import (
	"encoding/json"
	"net/http"
	"test/restserver/response"
	"test/usecases"
)

type RegisterUserHandler struct {
	uc usecases.UseCases
}

func (h RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in, err := createRegisterUserInputFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorMessage{err.Error()})
		return
	}

	out, err := h.uc.RegisterUser(in)
	if err != nil {
		status, resp := response.Error(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(registerUserResponse{
		User: struct {
			ID        string `json:"id"`
			Age       int    `json:"age"`
			IsMarried bool   `json:"is_married"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}{
			ID:        out.User.ID().String(),
			Age:       out.User.Age(),
			IsMarried: out.User.IsMarried(),
			Firstname: out.User.Firstname(),
			Lastname:  out.User.Lastname(),
		},
	})
}

func NewRegisterUser(uc usecases.UseCases) http.Handler {
	return RegisterUserHandler{
		uc: uc,
	}
}

func createRegisterUserInputFromRequest(r *http.Request) (usecases.RegisterUserInput, error) {
	var body registerUserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return usecases.RegisterUserInput{}, err
	}

	return usecases.RegisterUserInput{
		Age:       body.Age,
		IsMarried: body.IsMarried,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Password:  body.Password,
	}, nil
}

type registerUserRequest struct {
	Age       int    `json:"age"`
	IsMarried bool   `json:"is_married"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type registerUserResponse struct {
	User struct {
		ID        string `json:"id"`
		Age       int    `json:"age"`
		IsMarried bool   `json:"is_married"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"user"`
}

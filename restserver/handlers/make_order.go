package handlers

import (
	"encoding/json"
	"net/http"
	"test/domains"
	"test/restserver/response"
	"test/usecases"
)

type MakeOrderHandler struct {
	uc usecases.UseCases
}

func (h MakeOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in, err := makeOrderInputFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorMessage{err.Error()})
		return
	}

	out, err := h.uc.MakeOrder(in)
	if err != nil {
		status, resp := response.Error(err)
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createMakeOrderResponse{
		Order: struct {
			ID string `json:"id"`
		}{
			ID: out.Order.ID().String(),
		},
	})
}

func NewMakeOrder(uc usecases.UseCases) http.Handler {
	return RegisterUserHandler{
		uc: uc,
	}
}

func makeOrderInputFromRequest(r *http.Request) (usecases.MakeOrderInput, error) {
	var body createMakeOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return usecases.MakeOrderInput{}, err
	}

	return usecases.MakeOrderInput{
		User:     domains.User{},
		Products: []domains.ProductID{},
	}, nil
}

type createMakeOrderRequest struct {
	Products []int `json:"products"`
}

type createMakeOrderResponse struct {
	Order struct {
		ID string `json:"id"`
	} `json:"order"`
}

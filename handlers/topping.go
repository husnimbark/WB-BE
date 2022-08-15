package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "waysbucks/dto/result"
	toppingdto "waysbucks/dto/topping"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

var topping_file = "http://localhost:4000/uploads/"

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

// GET ALL
func (h *handlerTopping) FindToppings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	toppings, err := h.ToppingRepository.FindToppings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range toppings {
		toppings[i].Image = topping_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesFindToppings{Code: http.StatusOK, Data: toppings}
	json.NewEncoder(w).Encode(response)
}

// GET DETAIL
func (h *handlerTopping) GetTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var topping models.Topping
	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping.Image = topping_file + topping.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccesGetTopping{Code: http.StatusOK, Data: topping}
	json.NewEncoder(w).Encode(response)
}

// CREATE
func (h *handlerTopping) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET TOKEN
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") //GET THE VALUE
	filename := dataContex.(string)             //CONVERT TO STRING

	price, _ := strconv.Atoi(r.FormValue("price"))
	request := toppingdto.ToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping := models.Topping{
		Title:     request.Title,
		Price:     request.Price,
		Image:     filename,
		ProductID: userId,
	}

	topping, err = h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ = h.ToppingRepository.GetTopping(topping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccesGetTopping{Code: http.StatusOK, Data: topping}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTopping(u models.Topping) models.ToppingResponse {
	return models.ToppingResponse{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
	}
}

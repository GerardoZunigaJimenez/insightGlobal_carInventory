package handler

import (
	"encoding/json"
	"insightGlobal_carInventory/internal/model"
	"insightGlobal_carInventory/internal/service"
	"net/http"
	"strconv"

	"github.com/uptrace/bunrouter"
	"go.uber.org/zap"
)

type CarHandler interface {
	Create(w http.ResponseWriter, req bunrouter.Request) error
	GetByID(w http.ResponseWriter, req bunrouter.Request) error
	List(w http.ResponseWriter, req bunrouter.Request) error
	Update(w http.ResponseWriter, req bunrouter.Request) error
	Delete(w http.ResponseWriter, req bunrouter.Request) error
}

type carHandler struct {
	carService service.CarService
	log        *zap.SugaredLogger
}

func newCarHandler(carService service.CarService, log *zap.SugaredLogger) CarHandler {
	return &carHandler{
		carService: carService,
		log:        log,
	}
}

func (h carHandler) Create(w http.ResponseWriter, req bunrouter.Request) error {
	var car model.Car
	if err := json.NewDecoder(req.Body).Decode(&car); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, bunrouter.H{"message": "invalid request body"})
	}

	createdCar, err := h.carService.Create(req.Context(), &car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{"message": err.Error()})
	}

	w.WriteHeader(http.StatusCreated)
	return bunrouter.JSON(w, createdCar)
}

func (h carHandler) GetByID(w http.ResponseWriter, req bunrouter.Request) error {
	id := req.Params().ByName("id")

	car, err := h.carService.GetByID(req.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{"message": err.Error()})
	}
	if car == nil {
		w.WriteHeader(http.StatusNotFound)
		return bunrouter.JSON(w, bunrouter.H{"message": "car not found"})
	}

	return bunrouter.JSON(w, car)
}

func (h carHandler) List(w http.ResponseWriter, req bunrouter.Request) error {
	// Parse pageSize from query parameters, default to 10
	pageSize := 10
	if pageSizeStr := req.URL.Query().Get("pageSize"); pageSizeStr != "" {
		if parsedSize, err := strconv.Atoi(pageSizeStr); err == nil && parsedSize > 0 {
			pageSize = parsedSize
		}
	}

	// Parse pageNumber from query parameters, default to 1
	pageNumber := 1
	if pageNumberStr := req.URL.Query().Get("pageNumber"); pageNumberStr != "" {
		if parsedNumber, err := strconv.Atoi(pageNumberStr); err == nil && parsedNumber > 0 {
			pageNumber = parsedNumber
		}
	}

	cars, err := h.carService.List(req.Context(), pageSize, pageNumber)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{"message": err.Error()})
	}

	return bunrouter.JSON(w, cars)
}

func (h carHandler) Update(w http.ResponseWriter, req bunrouter.Request) error {
	var car model.Car
	if err := json.NewDecoder(req.Body).Decode(&car); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, bunrouter.H{"message": "invalid request body"})
	}

	updatedCar, err := h.carService.Update(req.Context(), &car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{"message": err.Error()})
	}

	return bunrouter.JSON(w, updatedCar)
}

func (h carHandler) Delete(w http.ResponseWriter, req bunrouter.Request) error {
	id := req.Params().ByName("id")

	if err := h.carService.Delete(req.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, bunrouter.H{"message": err.Error()})
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

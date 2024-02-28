package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request body"))
			w.Header().Set("Content-Type", "text/plain")
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// process
		v := internal.Vehicle{
			Id: reqBody.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}

		// - create vehicle
		err = h.sv.Create(v)
		if err != nil {
			// check error type with switch case and error.Is(err, internal.ErrVehicleAlreadyExists)
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				// response with status code 409 and message "vehicle already exists"
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("vehicle already exists"))
				w.Header().Set("Content-Type", "text/plain")
				response.JSON(w, http.StatusConflict, nil)
				return
			case errors.Is(err, internal.ErrVehicleIncomplete):
				// response with status code 400 and message "vehicle incomplete"
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("vehicle incomplete"))
				w.Header().Set("Content-Type", "text/plain")
				response.JSON(w, http.StatusBadRequest, nil)
				return
			default:
				// response with status code 500 and message "internal server error"
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				w.Header().Set("Content-Type", "text/plain")
				response.JSON(w, http.StatusInternalServerError, nil)
				return
			}
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/:color/:year
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting color and year from url params
		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")
		yearint, err := strconv.Atoi(year)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "year must be an integer")
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.GetByColorAndYear(color, yearint)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleColorEmpty), errors.Is(err, internal.ErrVehicleYearEmpty), errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByBrandAndYearRange is a method that returns a handler for the route - GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting brand, start_year and end_year from url params
		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")
		startYearint, err := strconv.Atoi(startYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "start_year must be an integer")
			return
		}
		endYearint, err := strconv.Atoi(endYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "end_year must be an integer")
			return
		}

		// process
		// - get vehicles by brand and year range
		v, err := h.sv.GetByBrandAndYearRange(brand, startYearint, endYearint)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleBrandEmpty), errors.Is(err, internal.ErrVehicleYearEmpty), errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			case errors.Is(err, internal.ErrVehicleYearEndInvalid):
				response.Error(w, http.StatusBadRequest, "end_year must be greater than start_year")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a handler for the route - GET /vehicles/average_speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting brand from url params
		brand := chi.URLParam(r, "brand")

		// process
		// - get average speed by brand
		v, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleBrandEmpty):
				response.Error(w, http.StatusBadRequest, "brand must not be empty")
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

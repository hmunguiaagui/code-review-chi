package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
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

// GetByBrandBetweenYears is a method that returns a handler for the route GET /vehicles/{brand}/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandBetweenYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get brand from path
		brand := r.URL.Query().Get("brand")
		// - get startYear from path
		startYear := r.URL.Query().Get("startYear")
		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("start_year must be an integer"))
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		// - get endYear from path
		endYear := r.URL.Query().Get("endYear")
		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("end_year must be an integer"))
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// process
		// - get vehicles by brand and years
		v, err := h.sv.FindByBrandBetweenYears(brand, startYearInt, endYearInt)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			// switch case for error types
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("vehicle not found"))
				response.JSON(w, http.StatusNotFound, nil)
			case errors.Is(err, internal.ErrVehicleBrandNotFound):
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("vehicle brand not found"))
				response.JSON(w, http.StatusNotFound, nil)
			case errors.Is(err, internal.ErrVehicleYearNotFound):
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("vehicle year not found"))
				response.JSON(w, http.StatusNotFound, nil)
			case errors.Is(err, internal.ErrVehicleYearInvalid):
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("vehicle year invalid"))
				response.JSON(w, http.StatusBadRequest, nil)
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
				response.JSON(w, http.StatusInternalServerError, nil)
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

package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

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

		// vehicle to vehicleJSON
		responseBody := VehicleJSON{
			ID:              v.Id,
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Height:          v.Height,
			Length:          v.Length,
			Width:           v.Width,
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    responseBody,
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

// CreateBatch is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting vehicles from request body
		var vehiclesJSON []VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&vehiclesJSON)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		// parse VehicleJSON to Vehicle
		vehicles := make([]internal.Vehicle, len(vehiclesJSON))
		for i, v := range vehiclesJSON {
			vehicles[i] = internal.Vehicle{
				Id: v.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           v.Brand,
					Model:           v.Model,
					Registration:    v.Registration,
					Color:           v.Color,
					FabricationYear: v.FabricationYear,
					Capacity:        v.Capacity,
					MaxSpeed:        v.MaxSpeed,
					FuelType:        v.FuelType,
					Transmission:    v.Transmission,
					Weight:          v.Weight,
					Dimensions: internal.Dimensions{
						Height: v.Height,
						Length: v.Length,
						Width:  v.Width,
					},
				},
			}
		}

		// - create vehicles
		v, err := h.sv.CreateBatch(vehicles)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleIncomplete):
				response.Error(w, http.StatusBadRequest, "vehicle/s incomplete")
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.Error(w, http.StatusConflict, "vehicle/s already exists")
			case errors.Is(err, internal.ErrVehicleBrandEmpty), errors.Is(err, internal.ErrVehicleYearEmpty), errors.Is(err, internal.ErrVehicleColorEmpty):
				response.Error(w, http.StatusBadRequest, "vehicle/s data must not be empty")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

// UpdateSpeedById is a method that returns a handler for the route - PUT /vehicles/{id}/update_speed?speed={speed}
func (h *VehicleDefault) UpdateSpeedById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting id and speed from url params
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "id must be an integer")
			return
		}
		speed := r.URL.Query().Get("speed")
		speedFloat, err := strconv.ParseFloat(speed, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "speed must be a number")
			return
		}

		// process
		// - update speed by id
		v, err := h.sv.UpdateSpeedById(idInt, speedFloat)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleIdInvalid):
				response.Error(w, http.StatusBadRequest, "invalid id")
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

// GetByFuelType is a method that returns a handler for the route - GET /vehicles/fuel_type/{type}
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting fuel type from url params
		fuelType := chi.URLParam(r, "type")

		// process
		// - get vehicles by fuel type
		v, err := h.sv.GetByFuelType(fuelType)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleFuelTypeEmpty):
				response.Error(w, http.StatusBadRequest, "fuel_type must not be empty")
			case errors.Is(err, internal.ErrVehicleNotFound):
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

// DeleteById is a method that returns a handler for the route - DELETE /vehicles/{id}
func (h *VehicleDefault) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting id from url params
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "id must be an integer")
			return
		}

		// process
		// - delete vehicle by id
		err = h.sv.DeleteById(idInt)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleIdInvalid):
				response.Error(w, http.StatusBadRequest, "invalid id")
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, nil)
	}
}

// GetByTransmissionType is a method that returns a handler for the route - GET /vehicles/transmission/{type}
func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting transmission type from url params
		transmissionType := chi.URLParam(r, "type")

		// process
		// - get vehicles by transmission type
		v, err := h.sv.GetByTransmissionType(transmissionType)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleTransmissionEmpty):
				response.Error(w, http.StatusBadRequest, "transmission_type must not be empty")
			case errors.Is(err, internal.ErrVehicleNotFound):
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

// UpdateFuelById is a method that returns a handler for the route - PUT /vehicles/{id}/update_fuel?fuel_type={fuel}
func (h *VehicleDefault) UpdateFuelById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting id and fuel type from url params
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "id must be an integer")
			return
		}
		fuelType := r.URL.Query().Get("fuel_type")

		// process
		// - update fuel type by id
		v, err := h.sv.UpdateFuelById(idInt, fuelType)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleIdInvalid):
				response.Error(w, http.StatusBadRequest, "invalid id")
			case errors.Is(err, internal.ErrVehicleFuelTypeEmpty):
				response.Error(w, http.StatusBadRequest, "fuel_type must not be empty")
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

// GetAverageCapacityByBrand is a method that returns a handler for the route - GET /vehicles/average_capacity/brand/{brand}
func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting brand from url params
		brand := chi.URLParam(r, "brand")

		// process
		// - get average capacity by brand
		v, err := h.sv.GetAverageCapacityByBrand(brand)
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

// GetByDimensions is a method that returns a handler for the route - GET /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting length and width from url params
		length := r.URL.Query().Get("length")
		// split length into min and max length
		lengthSplit := strings.Split(length, "-")
		// get min/max length from lengthSplit[0] and max length from lengthSplit[1] to float64
		minLengthFloat, err := strconv.ParseFloat(lengthSplit[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "min_length must be a float")
			return
		}
		maxLengthFloat, err := strconv.ParseFloat(lengthSplit[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "max_length must be a float")
			return
		}

		width := r.URL.Query().Get("width")
		// split width into min and max width
		widthSplit := strings.Split(width, "-")
		// get min/max width from widthSplit[0] and max width from widthSplit[1] to float64
		minWidthFloat, err := strconv.ParseFloat(widthSplit[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "min_width must be a float")
			return
		}
		maxWidthFloat, err := strconv.ParseFloat(widthSplit[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "max_width must be a float")
			return
		}

		// process
		// - get vehicles by dimensions
		v, err := h.sv.GetByDimensions(minLengthFloat, maxLengthFloat, minWidthFloat, maxWidthFloat)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleLengthInvalid), errors.Is(err, internal.ErrVehicleWidthInvalid):
				response.Error(w, http.StatusBadRequest, "invalid dimensions")
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make([]VehicleJSON, len(v))
		for i, value := range v {
			data[i] = VehicleJSON{
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
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByWeight is a method that returns a handler for the route - GET /vehicles/weight?min={weight_min}&max={weight_max}
func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// getting min and max weight from url params
		minWeight := r.URL.Query().Get("min")
		// convert min weight to float64
		minWeightFloat, err := strconv.ParseFloat(minWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "min_weight must be a number")
			return
		}
		maxWeight := r.URL.Query().Get("max")
		// convert max weight to float64
		maxWeightFloat, err := strconv.ParseFloat(maxWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "max_weight must be a number")
			return
		}

		// process
		// - get vehicles by weight
		v, err := h.sv.GetByWeight(minWeightFloat, maxWeightFloat)
		if err != nil {
			// switch to know the error type and response accordingly
			switch {
			case errors.Is(err, internal.ErrVehicleWeightInvalid):
				response.Error(w, http.StatusBadRequest, "invalid weight")
			case errors.Is(err, internal.ErrVehicleMinWeightGreaterThanMaxWeight):
				response.Error(w, http.StatusBadRequest, "min_weight must be less than or equal to max_weight")
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Error(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make([]VehicleJSON, len(v))
		for i, value := range v {
			data[i] = VehicleJSON{
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
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    data,
		})
	}
}

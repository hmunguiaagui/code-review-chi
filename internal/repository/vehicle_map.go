package repository

import (
	"app/internal"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a vehicle in the repository
func (r *VehicleMap) Create(v internal.Vehicle) (err error) {
	// check if vehicle already exists in db
	_, ok := r.db[v.Id]
	if ok {
		err = internal.ErrVehicleAlreadyExists
		return
	}
	// check if vehicle is incomplete
	if v.Id <= 0 || v.Brand == "" || v.Model == "" || v.Registration == "" || v.Color == "" || v.FabricationYear == 0 || v.Capacity == 0 || v.MaxSpeed == 0 || v.FuelType == "" || v.Transmission == "" || v.Weight == 0 || v.Height == 0 || v.Length == 0 || v.Width == 0 {
		err = internal.ErrVehicleIncomplete
		return
	}

	// add vehicle to db
	r.db[v.Id] = v

	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// check if color is empty
	if color == "" {
		err = internal.ErrVehicleColorEmpty
		return
	}

	// check if year is less or equal than zero
	if year <= 0 {
		err = internal.ErrVehicleYearEmpty
		return
	}

	// get vehicles by color and year
	for key, value := range r.db {
		if strings.EqualFold(value.Color, color) && value.FabricationYear == year {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (r *VehicleMap) GetByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// check if brand is empty
	if brand == "" {
		err = internal.ErrVehicleBrandEmpty
		return
	}

	// check if yearFrom is less or equal than zero
	if yearFrom <= 0 {
		err = internal.ErrVehicleYearEmpty
		return
	}

	// check if yearTo is less or equal than zero
	if yearTo <= 0 {
		err = internal.ErrVehicleYearEmpty
		return
	}

	// check if yearFrom is greater than yearTo
	if yearFrom > yearTo {
		err = internal.ErrVehicleYearEndInvalid
		return
	}

	// get vehicles by brand and year range
	for key, value := range r.db {
		if strings.EqualFold(value.Brand, brand) && value.FabricationYear >= yearFrom && value.FabricationYear <= yearTo {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

// GetAverageSpeedByBrand is a method that returns the average speed by brand
func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	// check if brand is empty
	if brand == "" {
		err = internal.ErrVehicleBrandEmpty
		return
	}

	// get vehicles by brand
	var vehicles map[int]internal.Vehicle = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if strings.EqualFold(value.Brand, brand) {
			vehicles[key] = value
		}
	}

	// check if map is empty
	if len(vehicles) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	// calculate average speed
	var totalSpeed float64
	for _, value := range vehicles {
		totalSpeed += value.MaxSpeed
	}
	averageSpeed = totalSpeed / float64(len(vehicles))

	return
}

// CreateBatch is a method that creates a batch of vehicles in the repository
func (r *VehicleMap) CreateBatch(v []internal.Vehicle) ([]internal.Vehicle, error) {
	// check if slice is empty
	if len(v) == 0 {
		return nil, internal.ErrVehicleBatchEmpty
	}

	// check if vehicle already exists in db
	for _, value := range v {
		_, ok := r.db[value.Id]
		if ok {
			return nil, internal.ErrVehicleAlreadyExists
		}
	}

	// check if slice has duplicates
	for i, value := range v {
		for j, value2 := range v {
			if i != j && value.Id == value2.Id {
				return nil, internal.ErrVehicleAlreadyExists
			}
		}
	}

	// check if slice has incomplete vehicles
	for _, value := range v {
		if value.Id <= 0 || value.Brand == "" || value.Model == "" || value.Registration == "" || value.Color == "" || value.FabricationYear == 0 || value.Capacity == 0 || value.MaxSpeed == 0 || value.FuelType == "" || value.Transmission == "" || value.Weight == 0 || value.Height == 0 || value.Length == 0 || value.Width == 0 {
			return nil, internal.ErrVehicleIncomplete
		}
	}

	// add vehicles to db
	for _, value := range v {
		r.db[value.Id] = value
	}

	return v, nil
}

// UpdateSpeedById is a method that updates the speed of a vehicle by id and returns the updated vehicle
func (r *VehicleMap) UpdateSpeedById(id int, speed float64) (vehicle internal.Vehicle, err error) {
	// check if id is less or equal than zero
	if id <= 0 {
		err = internal.ErrVehicleIdInvalid
		return
	}

	// check if speed is less or equal than zero
	if speed <= 0 {
		err = internal.ErrVehicleSpeedInvalid
		return
	}

	// check if vehicle exists in db
	vehicle, ok := r.db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	// update speed
	vehicle.MaxSpeed = speed

	// update vehicle in db
	r.db[id] = vehicle

	return
}

// GetByFuelType is a method that returns a map of vehicles by fuel type
func (r *VehicleMap) GetByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// check if fuel type is empty
	if fuelType == "" {
		err = internal.ErrVehicleFuelTypeEmpty
		return
	}

	// get vehicles by fuel type
	for key, value := range r.db {
		if strings.EqualFold(value.FuelType, fuelType) {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

// DeleteById is a method that deletes a vehicle by id
func (r *VehicleMap) DeleteById(id int) (err error) {
	// check if id is less or equal than zero
	if id <= 0 {
		err = internal.ErrVehicleIdInvalid
		return
	}

	// check if vehicle exists in db
	_, ok := r.db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	// delete vehicle from db
	delete(r.db, id)

	return
}

// GetByTransmissionType is a method that returns a map of vehicles by transmission type
func (r *VehicleMap) GetByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// check if transmission type is empty
	if transmissionType == "" {
		err = internal.ErrVehicleTransmissionEmpty
		return
	}

	// get vehicles by transmission type
	for key, value := range r.db {
		if strings.EqualFold(value.Transmission, transmissionType) {
			v[key] = value
		}
	}

	// check if map is empty
	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

// UpdateFuelById is a method that updates the fuel of a vehicle by id and returns the updated vehicle
func (r *VehicleMap) UpdateFuelById(id int, fuelType string) (vehicle internal.Vehicle, err error) {
	// check if id is less or equal than zero
	if id <= 0 {
		err = internal.ErrVehicleIdInvalid
		return
	}

	// check if fuel type is empty
	if fuelType == "" {
		err = internal.ErrVehicleFuelTypeEmpty
		return
	}

	// check if vehicle exists in db
	vehicle, ok := r.db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	// update fuel
	vehicle.FuelType = fuelType

	// update vehicle in db
	r.db[id] = vehicle

	return
}

// GetAverageCapacityByBrand is a method that returns the average capacity by brand
func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	// check if brand is empty
	if brand == "" {
		err = internal.ErrVehicleBrandEmpty
		return
	}

	// get vehicles by brand
	var vehicles map[int]internal.Vehicle = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if strings.EqualFold(value.Brand, brand) {
			vehicles[key] = value
		}
	}

	// check if map is empty
	if len(vehicles) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	// calculate average capacity
	var totalCapacity float64
	for _, value := range vehicles {
		totalCapacity += float64(value.Capacity)
	}
	averageCapacity = totalCapacity / float64(len(vehicles))

	return
}

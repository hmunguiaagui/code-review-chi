package repository

import "app/internal"

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

// FindByBrandBetweenYears is a method that returns a map of vehicles that match the brand and years
func (r *VehicleMap) FindByBrandBetweenYears(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// check if brand is empty
	if brand == "" {
		err = internal.ErrVehicleBrandNotFound
		return
	}

	// check if brand exists
	found := false
	for _, value := range r.db {
		if value.Brand == brand {
			found = true
			break
		}
	}

	if !found {
		err = internal.ErrVehicleBrandNotFound
		return
	}

	// check if startYear and endYear are valid
	if startYear <= 0 || endYear <= 0 {

		err = internal.ErrVehicleYearInvalid
		return
	}

	// check if startYear is less than endYear
	if startYear > endYear {
		err = internal.ErrVehicleYearNotFound
		return
	}

	// copy db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	return
}

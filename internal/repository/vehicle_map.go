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

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// Switch case for color and year empty or not empty
	switch {
	case color == "" && year != 0:
		// Get vehicles by year
		for key, value := range r.db {
			if value.FabricationYear == year {
				v[key] = value
			}
		}
	case year == 0 && color != "":
		// Get vehicles by color
		for key, value := range r.db {
			if value.Color == color {
				v[key] = value
			}
		}
	default:
		// Get vehicles by color and year
		for key, value := range r.db {
			if value.Color == color && value.FabricationYear == year {
				v[key] = value
			}
		}
	}

	// check if v is empty
	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	return
}

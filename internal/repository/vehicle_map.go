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

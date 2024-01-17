package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	return
}

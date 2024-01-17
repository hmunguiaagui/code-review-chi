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

// FindByBrandBetweenYears is a method that returns a map of vehicles that match the brand and years
func (s *VehicleDefault) FindByBrandBetweenYears(brand string, startYear, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandBetweenYears(brand, startYear, endYear)
	return
}

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

// Create is a method that creates a vehicle in the repository
func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	return
}

// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (s *VehicleDefault) GetByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByBrandAndYearRange(brand, yearFrom, yearTo)
	return
}

// GetAverageSpeedByBrand is a method that returns the average speed by brand
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)
	return
}

// CreateBatch is a method that creates a batch of vehicles
func (s *VehicleDefault) CreateBatch(v []internal.Vehicle) ([]internal.Vehicle, error) {
	vehicles, err := s.rp.CreateBatch(v)
	return vehicles, err
}

// UpdateSpeedById is a method that updates the speed of a vehicle by id and returns the vehicle updated
func (s *VehicleDefault) UpdateSpeedById(id int, speed float64) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.UpdateSpeedById(id, speed)
	return
}

// GetByFuelType is a method that returns a map of vehicles by fuel type
func (s *VehicleDefault) GetByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(fuelType)
	return
}

// DeleteById is a method that deletes a vehicle by id
func (s *VehicleDefault) DeleteById(id int) (err error) {
	err = s.rp.DeleteById(id)
	return
}

// GetByTransmissionType is a method that returns a map of vehicles by transmission type
func (s *VehicleDefault) GetByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByTransmissionType(transmissionType)
	return
}

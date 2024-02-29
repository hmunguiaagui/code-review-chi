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

// FindAll is a method that returns a slice of all vehicles
func (s *VehicleDefault) FindAll() ([]internal.Vehicle, error) {
	return s.rp.FindAll()
}

// Create is a method that creates a vehicle in the repository
func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// GetByColorAndYear is a method that returns a slice of vehicles by color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) ([]internal.Vehicle, error) {
	return s.rp.GetByColorAndYear(color, year)
}

// GetByBrandAndYearRange is a method that returns a slice of vehicles by brand and year range
func (s *VehicleDefault) GetByBrandAndYearRange(brand string, yearFrom int, yearTo int) ([]internal.Vehicle, error) {
	return s.rp.GetByBrandAndYearRange(brand, yearFrom, yearTo)
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

// GetByFuelType is a method that returns a slice of vehicles by fuel type
func (s *VehicleDefault) GetByFuelType(fuelType string) ([]internal.Vehicle, error) {
	return s.rp.GetByFuelType(fuelType)
}

// DeleteById is a method that deletes a vehicle by id
func (s *VehicleDefault) DeleteById(id int) (err error) {
	err = s.rp.DeleteById(id)
	return
}

// GetByTransmissionType is a method that returns a slice of vehicles by transmission type
func (s *VehicleDefault) GetByTransmissionType(transmissionType string) ([]internal.Vehicle, error) {
	return s.rp.GetByTransmissionType(transmissionType)
}

// UpdateFuelById is a method that updates the fuel of a vehicle by id and returns the vehicle updated
func (s *VehicleDefault) UpdateFuelById(id int, fuel string) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.UpdateFuelById(id, fuel)
	return
}

// GetAverageCapacityByBrand is a method that returns the average capacity by brand
func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	averageCapacity, err = s.rp.GetAverageCapacityByBrand(brand)
	return
}

// GetByDimensions is a method that returns a slice of vehicles by dimensions (min length, max length, min width, max width)
func (s *VehicleDefault) GetByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v []internal.Vehicle, err error) {
	v, err = s.rp.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
	return
}

// GetByWeight is a method that returns a slice of vehicles by weight (min weight, max weight)
func (s *VehicleDefault) GetByWeight(minWeight float64, maxWeight float64) (v []internal.Vehicle, err error) {
	v, err = s.rp.GetByWeight(minWeight, maxWeight)
	return
}

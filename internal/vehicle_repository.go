package internal

import "errors"

var (
	// ErrVehicleNotFound is an error that represents a vehicle not found
	ErrVehicleNotFound = errors.New("vehicle not found")
	// ErrVehicleBrandNotFound is an error that represents a vehicle brand not found
	ErrVehicleBrandNotFound = errors.New("vehicle brand not found")
	// ErrVehicleYearNotFound is an error that represents a vehicle year not found
	ErrVehicleYearNotFound = errors.New("vehicle year not found")
	// ErrVehicleYearInvalid is an error that represents a vehicle year invalid
	ErrVehicleYearInvalid = errors.New("vehicle year invalid")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindByBrandBetweenYears is a method that returns a map of vehicles that match the brand and years
	FindByBrandBetweenYears(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
}

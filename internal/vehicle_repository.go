package internal

import "errors"

var (
	// ErrVehicleNotFound is an error that represents a vehicle not found
	ErrVehicleNotFound = errors.New("vehicle not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindByColorAndYear is a method that returns a map of vehicles by color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
}

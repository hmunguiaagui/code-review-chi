package internal

import "errors"

var (
	// ErrVehicleAlreadyExists is an error that represents that the vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	// ErrVehicleIncomplete is an error that represents that the vehicle is incomplete
	ErrVehicleIncomplete = errors.New("vehicle is incomplete")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// Create is a method that creates a vehicle in the repository
	Create(v Vehicle) (err error)
}

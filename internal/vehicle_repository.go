package internal

import "errors"

var (
	// ErrVehicleAlreadyExists is an error that represents that the vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	// ErrVehicleIncomplete is an error that represents that the vehicle is incomplete
	ErrVehicleIncomplete = errors.New("vehicle is incomplete")
	// ErrVehicleColorEmpty is an error that represents that the color is empty
	ErrVehicleColorEmpty = errors.New("color is empty")
	// ErrVehicleYearEmpty is an error that represents that the year is empty
	ErrVehicleYearEmpty = errors.New("year is empty")
	// ErrVehicleYearEndInvalid is an error that represents that the year end is invalid
	ErrVehicleYearEndInvalid = errors.New("year end is invalid")
	// ErrVehicleNotFound is an error that represents that the vehicle was not found
	ErrVehicleNotFound = errors.New("vehicle not found")
	// ErrVehicleBrandEmpty is an error that represents that the brand is empty
	ErrVehicleBrandEmpty = errors.New("brand is empty")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a vehicle in the repository
	Create(v Vehicle) (err error)
	// GetByColorAndYear is a method that returns a map of vehicles by color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
	GetByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]Vehicle, err error)
	// GetAverageSpeedByBrand is a method that returns the average speed by brand
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
}

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
	// ErrVehicleBatchEmpty	is an error that represents that the batch is empty
	ErrVehicleBatchEmpty = errors.New("batch is empty")
	// ErrVehicleIdInvalid is an error that represents that the id is invalid
	ErrVehicleIdInvalid = errors.New("id is invalid")
	// ErrVehicleSpeedInvalid is an error that represents that the speed is invalid
	ErrVehicleSpeedInvalid = errors.New("speed is invalid")
	// ErrVehicleFuelTypeEmpty is an error that represents that the fuel type is empty
	ErrVehicleFuelTypeEmpty = errors.New("fuel type is empty")
	// ErrVehicleTransmissionEmpty is an error that represents that the transmission is empty
	ErrVehicleTransmissionEmpty = errors.New("transmission is empty")
	// ErrVehicleLengthInvalid is an error that represents that the length is invalid
	ErrVehicleLengthInvalid = errors.New("length is invalid")
	// ErrVehicleWidthInvalid is an error that represents that the width is invalid
	ErrVehicleWidthInvalid = errors.New("width is invalid")
	// ErrVehicleMinLengthGreaterThanMaxLength is an error that represents that the min length is greater than the max length
	ErrVehicleMinLengthGreaterThanMaxLength = errors.New("min length is greater than max length")
	// ErrVehicleMinWidthGreaterThanMaxWidth is an error that represents that the min width is greater than the max width
	ErrVehicleMinWidthGreaterThanMaxWidth = errors.New("min width is greater than max width")
	// ErrVehicleWeightInvalid is an error that represents that the weight is invalid
	ErrVehicleWeightInvalid = errors.New("weight is invalid")
	// ErrVehicleMinWeightGreaterThanMaxWeight is an error that represents that the min weight is greater than the max weight
	ErrVehicleMinWeightGreaterThanMaxWeight = errors.New("min weight is greater than max weight")
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
	// CreateBatch is a method that creates a batch of vehicles in the repository
	CreateBatch(v []Vehicle) ([]Vehicle, error)
	// UpdateSpeedById is a method that updates the speed of a vehicle by id and returns the vehicle updated
	UpdateSpeedById(id int, speed float64) (vehicle Vehicle, err error)
	// GetByFuelType is a method that returns a map of vehicles by fuel type
	GetByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// DeleteById is a method that deletes a vehicle by id
	DeleteById(id int) (err error)
	// GetByTransmissionType is a method that returns a map of vehicles by transmission type
	GetByTransmissionType(transmissionType string) (v map[int]Vehicle, err error)
	// UpdateFuelById is a method that updates the fuel of a vehicle by id and returns the vehicle updated
	UpdateFuelById(id int, fuelType string) (vehicle Vehicle, err error)
	// GetAverageCapacityByBrand is a method that returns the average capacity by brand
	GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error)
	// GetByDimensions is a method that returns a slice of vehicles by dimensions (min length, max length, min width, max width)
	GetByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v []Vehicle, err error)
	// GetByWeight is a method that returns a slice of vehicles by weight (min weight, max weight)
	GetByWeight(minWeight float64, maxWeight float64) (v []Vehicle, err error)
}

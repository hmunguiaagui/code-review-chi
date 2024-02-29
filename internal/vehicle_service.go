package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
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
	UpdateFuelById(id int, fuel string) (vehicle Vehicle, err error)
	// GetAverageCapacityByBrand is a method that returns the average capacity by brand
	GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error)
	// GetByDimensions is a method that returns a slice of vehicles by dimensions (min length, max length, min width, max width)
	GetByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v []Vehicle, err error)
}

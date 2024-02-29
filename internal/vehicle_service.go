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
}

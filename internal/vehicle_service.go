package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindByBrandBetweenYears is a method that returns a map of vehicles that match the brand and years
	FindByBrandBetweenYears(brand string, startYear, endYear int) (v map[int]Vehicle, err error)
}

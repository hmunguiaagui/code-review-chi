package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindByColorAndYear is a method that returns a map of vehicles by color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
}

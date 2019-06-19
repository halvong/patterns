package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSets() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//Product
type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

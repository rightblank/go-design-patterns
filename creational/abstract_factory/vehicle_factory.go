package abstract_factory

import "errors"

type VehicleFactory interface {
	GetVehicle(t int) (Vehicle, error)
}

const (
	CarFactoryType = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New("")
	}
}
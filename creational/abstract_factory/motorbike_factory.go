package abstract_factory

import "errors"

const SportMotorbikeType =1
const CruiseMotorbikeType  = 2

type MotorbikeFactory struct {

}

func (*MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(LuxuryCar), nil
	case CruiseMotorbikeType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("motorbike")


	}
}
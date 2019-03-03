package abstract_factory

type CruiseMotorbike struct {

}

func (l *CruiseMotorbike) GetType() int{
	return CruiseMotorbikeType
}

func (l *CruiseMotorbike) GetWheels() int {
	return 4
}

func (l *CruiseMotorbike) GetSeats() int {
	return 5
}

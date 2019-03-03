package abstract_factory

type SportMotorbike struct {

}

func (l *SportMotorbike) GetType() int{
	return SportMotorbikeType
}

func (l *SportMotorbike) GetWheels() int {
	return 4
}

func (l *SportMotorbike) GetSeats() int {
	return 5
}

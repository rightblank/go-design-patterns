package creational

type singleton struct {
	count int
}

var instance *singleton

//GetInstance dd
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (instance *singleton) AddOne() int {
	instance.count++
	return instance.count
}

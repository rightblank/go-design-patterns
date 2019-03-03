package creational

type BuildProcess interface {
	setWheels() BuildProcess
	setSeats() BuildProcess
	setStructure() BuildProcess
	GetProduct() Vehicle
}

type Vehicle struct {
	Wheels, Seats int
	Structure     string
}


// could be constructed as a singleton
type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() BuildProcess{
	if m.builder == nil {
		panic("builder not exist")
	}
	m.builder.setSeats().setWheels().setStructure()
	return m.builder
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

//////=========================
type CarBuilder struct {
	v Vehicle
}

func (b *CarBuilder) setWheels() BuildProcess {
	b.v.Wheels = 4
	return b
}

func (b *CarBuilder) setSeats() BuildProcess {
	b.v.Seats = 5
	return b
}

func (b *CarBuilder) setStructure() BuildProcess {
	b.v.Structure = "Car"
	return b
}

func (b *CarBuilder) GetProduct() Vehicle {
	return b.v
}

//====================================================
type BusBuilder struct {
	v Vehicle
}

func (b *BusBuilder) setWheels() BuildProcess {
	b.v.Wheels = 4
	return b
}

func (b *BusBuilder) setSeats() BuildProcess {
	b.v.Seats = 5
	return b
}

func (b *BusBuilder) setStructure() BuildProcess {
	b.v.Structure = "Car"
	return b
}

func (b *BusBuilder) GetProduct() Vehicle {
	return b.v
}
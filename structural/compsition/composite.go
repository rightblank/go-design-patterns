package compsition

type Athlete struct {
}

func (a *Athlete) Train() {
    println("Training")
}

func Swim() {
    println("Swimming")
}

type CompositeSwimmerA struct {
    MyAthlete Athlete
    MySwim    *func()
}

//---------------------------------------
type Trainer interface {
    Train()
}

type Swimmer interface {
    Swim()
}

type SwimmerImplementor struct {
}

func (s *SwimmerImplementor) Swim() {
    println("Swimming")
}

type CompositeSwimmerB struct {
    Trainer
    Swimmer
}

//-------------------------------------
type Animal struct {
}

func (a *Animal) Eat() {
    println("Eating")
}

type Shark struct {
    Animal
    Swim func()
}

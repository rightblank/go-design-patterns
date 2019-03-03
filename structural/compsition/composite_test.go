package compsition

import "testing"

func TestAthlete_Train(t *testing.T) {
    athlete := Athlete{}
    athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
    localSwim := Swim
    swimmer := CompositeSwimmerA{
        MySwim: &localSwim,
    }
    swimmer.MyAthlete.Train()
    (*swimmer.MySwim)()
}

func TestAnimal_Eat(t *testing.T) {
    shark := Shark{
        Swim: Swim,
    }

    shark.Eat()
    shark.Swim()
}

func TestSwim2(t *testing.T) {
    swimmer := CompositeSwimmerB{
        &Athlete{},
        &SwimmerImplementor{},
    }

    swimmer.Swim()
    swimmer.Train()
}

package flyweight

import (
    "fmt"
    "testing"
)

func TestTeamFlyWeightFactory_GetTeam(t *testing.T) {
    factory := NewTeamFactory()

    teamA := factory.GetTeam(TEAM_A)
    if teamA == nil {
        t.Error("the pointer to team_a is nil")
    }

    teamA2 := factory.GetTeam(TEAM_A)
    if teamA2 == nil {
        t.Error("the pointer to team_a is nil")
    }

    if teamA != teamA2 {
        t.Error("team a objects isn't the same")
    }
    if factory.GetNumberOfObjects() != 1 {
        t.Errorf("the number of objects created is not 1: %d", factory.GetNumberOfObjects())
    }
}

func Test_HighVolume(t *testing.T) {
    factory := NewTeamFactory()

    teams := make([]*Team, 500000*2)
    for i := 0; i < 500000; i++ {
        teams[i] = factory.GetTeam(TEAM_A)
    }
    for i := 0; i < 500000; i++ {
        teams[i] = factory.GetTeam(TEAM_B)
    }

    if factory.GetNumberOfObjects() != 2 {
        t.Errorf("the number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
    }

    for i := 0; i < 3; i++ {
        fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
    }
}

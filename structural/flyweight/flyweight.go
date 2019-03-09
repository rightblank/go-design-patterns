package flyweight

import "time"

const (
    TEAM_A = iota
    TEAM_B
)

type Player struct {
    Name         string
    Surname      string
    PreviousTeam uint64
    Photo        []byte
}

type HistoricalData struct {
    Year          uint8
    LeagueResults []Match
}

type Team struct {
    ID             uint64
    Name           string
    Shield         []byte
    Playes         []Player
    HistoricalData []HistoricalData
}

type Match struct {
    Date        time.Time
    VisitorID   uint64
    LocalStore  byte
    VisitScore  byte
    LocalShoots uint16
    VisitShoots uint16
}

func getTeamFacotry(team int) Team {
    switch team {

    case TEAM_B:
        return Team{
            ID:   2,
            Name: "TEAM_B",
        }
    default:
        return Team{
            ID:   2,
            Name: "TEAM_A",
        }
    }

}

type teamFlyWeightFactory struct {
    createdTeams map[int]*Team
}

func NewTeamFactory() teamFlyWeightFactory {
    return teamFlyWeightFactory{
        createdTeams: make(map[int]*Team, 0),
    }
}

func (t *teamFlyWeightFactory) GetTeam(teamName int) *Team {
    if nil != t.createdTeams[teamName] {
        return t.createdTeams[teamName]
    }

    team := getTeamFacotry(teamName)
    t.createdTeams[teamName] = &team
    return t.createdTeams[teamName]
}

func (t *teamFlyWeightFactory) GetNumberOfObjects() int {
    return len(t.createdTeams)
}

package domain

import (
	"awesomeProject/DoomParser/cmd/misc"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//go:generate mockgen -source=./match.go -destination=./mock/usecase_mock.go -package=domain
type MatchUsecase interface {
	SplitByMatch([]string) ([]Match, error)
	ParseKills(match Match) (Kills, int, error)
}

type Match struct {
	ID         uint64
	TotalKills int
	Players    []string
	Kills      Kills
	Commands   []CommandInfo
}

type CommandInfo struct {
	Time            string
	CommandName     string
	UnparsedCommand string
}

type Kills map[string]int

func (m *Match) AddKill(killer string) {
	if _, ok := m.Kills[killer]; ok {
		m.Kills[killer]++
	} else {
		m.Kills = make(Kills)
		m.Kills[killer] = 1
	}
}

func (m *Match) AddWorldKill(killed string) {
	if _, ok := m.Kills[killed]; ok {
		m.Kills[killed]--
	} else {
		m.Kills[killed] = -1
	}
}

func (m *Match) InsertPlayer(player string) {
	for _, playerExistant := range m.Players {
		if playerExistant == player {
			return
		}
	}
	m.Players = append(m.Players, player)
	return
}

func (m *Match) ParseKills() error {
	for _, command := range m.Commands {
		if command.CommandName == misc.Kill {
			m.TotalKills++
			killer, killed, err := GetKillInfo(command.UnparsedCommand)
			fmt.Println(killer, killed)
			if err == nil {
				if strings.Contains(killer, "<world>") {
					m.AddWorldKill(killed)
				} else {
					m.InsertPlayer(killer)
					m.AddKill(killer)
				}
				m.InsertPlayer(killed)
			}

		}
	}
	return nil
}

func GetKillInfo(unparsedCommand string) (string, string, error) {
	var (
		killer string
		killed string
	)
	regex := regexp.MustCompile(`:(.*?)killed(.*?)by`)
	matches := regex.FindStringSubmatch(unparsedCommand)
	if len(matches) == 3 {
		killer = matches[1]
		killed = matches[2]
		return killer, killed, nil
	}
	return "", "", errors.New(misc.ErrNotACommand)
}

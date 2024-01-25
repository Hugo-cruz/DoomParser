package usecase

import (
	"awesomeProject/DoomParser/cmd/misc"
	"awesomeProject/DoomParser/cmd/parser/domain"
	"errors"
	"strings"
)

type matchUsecase struct {
	logRepository domain.LogRepository
}

func NewMatchUsecase(logRepository domain.LogRepository) domain.MatchUsecase {
	return &matchUsecase{logRepository: logRepository}
}

func (m *matchUsecase) ParseKills(match domain.Match) (domain.Kills, int, error) {
	err := match.ParseKills()
	if err != nil {
		return nil, 0, err
	}
	return match.Kills, match.TotalKills, nil
}

func (m *matchUsecase) SplitByMatch(logs []string) ([]domain.Match, error) {
	var (
		match domain.Match
	)

	if len(logs) == 0 {
		return nil, errors.New(misc.ErrEmptyLogs)
	}
	matches := make([]domain.Match, 0)

	sanitizedLogs := sanitizeLogs(logs)

	for _, logLine := range sanitizedLogs {
		time, instruction, commandInfo := splitLine(logLine)
		command := domain.CommandInfo{
			Time:            time,
			CommandName:     instruction,
			UnparsedCommand: commandInfo,
		}
		match.Commands = append(match.Commands, command)
		if instruction == misc.ShutdownGame || (instruction == misc.InitGame && len(match.Commands) < 1) {
			match.Kills = make(domain.Kills)
			matches = append(matches, match)
			match = domain.Match{}
		}

	}
	return matches, nil
}

func sanitizeLogs(logs []string) []string {
	var sanitizedLogs []string

	for _, log := range logs {
		if isACommand(log) {
			sanitizedLogs = append(sanitizedLogs, log)
		}

	}
	return sanitizedLogs
}

func splitLine(line string) (string, string, string) {
	info := strings.Split(line, " ")
	info = removeEmptyStrings(info)
	time := info[0]
	command := info[1]
	commandInfo := strings.Join(info[2:], " ")
	return time, command, commandInfo
}

func isACommand(command string) bool {
	return !strings.Contains(command, misc.NotACommand)
}

func removeEmptyStrings(elements []string) []string {
	var result []string

	for _, element := range elements {
		if strings.TrimSpace(element) != "" {
			result = append(result, element)
		}
	}

	return result
}

package repository

import (
	"awesomeProject/DoomParser/cmd/parser/domain"
	"bufio"
	"os"
)

type logFileRepository struct {
	fileName string
	file     *os.File
}

func NewLogFileRepository(fileName string) (domain.LogRepository, error) {
	return &logFileRepository{fileName: fileName}, nil
}

func (l *logFileRepository) GetAll() ([]string, error) {
	file, err := os.Open(l.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

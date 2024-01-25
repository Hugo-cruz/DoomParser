package usecase

import "awesomeProject/DoomParser/cmd/parser/domain"

type fileUsecase struct {
	logRepository domain.LogRepository
}

func NewFileUsecase(logRepository domain.LogRepository) domain.FileUsecase {
	return &fileUsecase{logRepository: logRepository}
}

func (u *fileUsecase) GetAll() ([]string, error) {
	return u.logRepository.GetAll()
}

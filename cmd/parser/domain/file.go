package domain

//go:generate mockgen -source=./file.go -destination=./mock/repository_mock.go -package=domain
type LogRepository interface {
	GetAll() ([]string, error)
}

type FileUsecase interface {
	GetAll() ([]string, error)
}

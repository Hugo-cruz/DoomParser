package parser

import (
	"awesomeProject/DoomParser/cmd/misc"
	"awesomeProject/DoomParser/cmd/parser/domain"
	"awesomeProject/DoomParser/cmd/parser/repository"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Dependencies struct {
	fileRepository domain.LogRepository
}

type Config struct {
	filePath string
}

func BuildDependencies(config Config) Dependencies {
	var (
		err            error
		fileRepository domain.LogRepository
	)

	if fileRepository, err = repository.NewLogFileRepository(config.filePath); err != nil {
		fmt.Println("Error opening repository")
		fmt.Println(err)
		panic(err)
	}
	return Dependencies{
		fileRepository: fileRepository,
	}

}

func ReadConfig(configFilePath string) (Config, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	gopath, err := getGoHome()
	if err != nil {
		return Config{}, err

	}
	config.filePath = gopath + config.filePath
	return config, nil
}

func getGoHome() (string, error) {
	gohome := os.Getenv("GOHOME")
	if gohome == "" {
		fmt.Println(misc.ErrGoHome)
		return "", errors.New(misc.ErrGoHome)
	} else {
		return gohome, nil
	}
}

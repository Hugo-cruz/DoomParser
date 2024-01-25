package parser

import (
	"awesomeProject/DoomParser/cmd/parser/usecase"
	"fmt"
)

func RunApplication(dependencies Dependencies) {

	fileUseCase := usecase.NewFileUsecase(dependencies.fileRepository)
	matchUseCase := usecase.NewMatchUsecase(dependencies.fileRepository)

	logLines, err := fileUseCase.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	matches, err := matchUseCase.SplitByMatch(logLines)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, match := range matches {
		err = match.ParseKills()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(match.ID)
		fmt.Println(match.TotalKills)
		fmt.Println(match.Players)
		fmt.Println("Kills")
		for killerName, numberOfKills := range match.Kills {
			fmt.Println(killerName, numberOfKills)
		}
		fmt.Println("--------------------")
	}

}

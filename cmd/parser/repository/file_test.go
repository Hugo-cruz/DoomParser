package repository

import (
	domain2 "awesomeProject/DoomParser/cmd/parser/domain"
	"awesomeProject/DoomParser/cmd/parser/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
	mockRepository *domain.MockLogRepository
	repository     domain2.LogRepository
}

func (suite *FileTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.mockRepository = domain.NewMockLogRepository(ctrl)
	suite.repository, _ = NewLogFileRepository("teste")
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}

func (suite *FileTestSuite) TestNewLogFileRepository() {
	suite.T().Run("should open file without fail", func(t *testing.T) {

		oldFs := domain.NewFs
		mfs := &domain.MockedFS{}
		domain.NewFs = mfs
		defer func() {
			domain.NewFs = oldFs
		}()

		_, err := suite.repository.GetAll()
		assert.NotNil(t, err)

	})
}

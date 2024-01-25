package usecase

import (
	domain "awesomeProject/DoomParser/cmd/parser/domain/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestFileSuite struct {
	suite.Suite
	repository *domain.MockLogRepository
}

func (suite *TestFileSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.repository = domain.NewMockLogRepository(ctrl)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestFileSuite))
}

func (suite *TestFileSuite) TestSuccess() {

	suite.Run("should get all info with success", func() {
		suite.repository.EXPECT().GetAll().Return([]string{"test"}, nil)
		_, err := suite.repository.GetAll()
		assert.Nil(suite.T(), err)
	})
	suite.Run("should throw error", func() {
		suite.repository.EXPECT().GetAll().Return([]string{"test"}, errors.New("error"))
		_, err := suite.repository.GetAll()
		assert.NotNil(suite.T(), err)
	})
}

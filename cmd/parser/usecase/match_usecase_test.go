package usecase

import (
	"awesomeProject/DoomParser/cmd/misc"
	domain2 "awesomeProject/DoomParser/cmd/parser/domain"
	domain "awesomeProject/DoomParser/cmd/parser/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestMatchSuite struct {
	suite.Suite
	MockRepository *domain.MockLogRepository
	UseCase        domain2.MatchUsecase
}

func (suite *TestMatchSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.MockRepository = domain.NewMockLogRepository(ctrl)
	suite.UseCase = NewMatchUsecase(suite.MockRepository)
}

func TestRunSuiteMatch(t *testing.T) {
	suite.Run(t, new(TestMatchSuite))
}

func (suite *TestMatchSuite) TestSuccess() {
	suite.Run("should get all info with success", func() {

		matchInfo, err := suite.UseCase.SplitByMatch(mockLogInfo())
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), matchInfo)
	})
	suite.Run("should return the correct size of number of matches", func() {
		suite.MockRepository.EXPECT().GetAll().Return(mockLogInfo(), nil)

		matchInfo, err := suite.UseCase.SplitByMatch(mockLogInfo())
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), len(matchInfo), 1)
	})
	suite.Run("should return the correct kill object", func() {
		suite.MockRepository.EXPECT().GetAll().Return(mockLogInfo(), nil)

		matchInfo, err := suite.UseCase.SplitByMatch(mockLogInfo())
		kills, _, err := suite.UseCase.ParseKills(matchInfo[0])
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), 1, len(kills))
	})
	suite.Run("should return the correct total kills", func() {
		suite.MockRepository.EXPECT().GetAll().Return(mockLogInfo(), nil)

		matchInfo, _ := suite.UseCase.SplitByMatch(mockLogInfo())
		_, totalKills, err := suite.UseCase.ParseKills(matchInfo[0])
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), totalKills, 5)
	})

}

func (suite *TestMatchSuite) TestFail() {
	suite.Run("should throw error when log file is empty", func() {
		suite.MockRepository.EXPECT().GetAll().Return([]string{}, nil)

		_, err := suite.UseCase.SplitByMatch([]string{})
		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), misc.ErrEmptyLogs, err.Error())
	})

}

func mockLogInfo() []string {
	mockInfo := []string{
		`0:00 ------------------------------------------------------------`,
		`0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`,
		`15:00 Exit: Timelimit hit.`,
		`20:34 ClientConnect: 2`,
		`20:34 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0`,
		`20:37 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0`,
		`20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`,
		`21:07 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`,
		`21:07 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH`,
		`21:07 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH`,
		`22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH`,
		`20:37 ClientBegin: 2`,
		`20:37 ShutdownGame:`,
		`20:37 ------------------------------------------------------------`,
		`20:37 ------------------------------------------------------------`,
	}
	return mockInfo
}

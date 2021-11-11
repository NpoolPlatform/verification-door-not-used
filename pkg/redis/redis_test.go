package redis

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestRedis(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	client := Client()
	assert.NotNil(t, client)

	userID := uuid.New().String()
	verifyCode := VerifyUserCode{
		Code:     "123456",
		SendTime: time.Now().Unix(),
	}

	err := InsertKeyInfo(userID, "test", verifyCode, 0)
	fmt.Println("test error is:", err)
	assert.Nil(t, err)

	info, err := QueryVerifyCodeKeyInfo(userID, "test")
	if assert.Nil(t, err) {
		assert.Equal(t, info.Code, verifyCode.Code)
		assert.Equal(t, info.SendTime, verifyCode.SendTime)
	}
}
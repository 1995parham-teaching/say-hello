package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cng-by-example/say-hello/internal/handler"
	"github.com/cng-by-example/say-hello/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type HelloSuite struct {
	suite.Suite

	engine *fiber.App
}

func (suite *HelloSuite) SetupSuite() {
	suite.engine = fiber.New()

	handler.NewHello().Register(suite.engine.Group(""))
}

func (suite *HelloSuite) TestHandler() {
	require := suite.Require()

	req := httptest.NewRequest("GET", "/hello", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	resp, err := suite.engine.Test(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)

	msg := new(response.Message)
	require.NoError(json.NewDecoder(resp.Body).Decode(msg))
	require.True(strings.HasPrefix(msg.Msg, "Hello at"))

	require.NoError(resp.Body.Close())
}

func TestHelloSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HelloSuite))
}

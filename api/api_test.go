package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ken-aio/echo-base/db"
	"github.com/ken-aio/echo-base/middleware"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	. "gopkg.in/check.v1"
)

type ApiSuite struct{}

func (s *ApiSuite) SetUpSuite(c *C) {
	// テストスイート全体での初期化処理を記述する
}

func (s *ApiSuite) SetUpTest(c *C) {
	// 個々のテストごとに実行される初期化時の共通処理を記述する
}

func (s *ApiSuite) TearDownSuite(c *C) {
	// テストスイート全体での終了後処理を記述する
}

func (s *ApiSuite) TearDownTest(c *C) {
	// 個々のテストごとに実行される終了後の共通処理を記述する
}

func buildContext(c *C, method string, url string, params string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req, err := http.NewRequest(method, url, strings.NewReader(params))
	c.Assert(err, IsNil)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	return context, rec
}

func requestAPI(c *C, handler echo.HandlerFunc, context echo.Context, rec *httptest.ResponseRecorder) (*dbr.Tx, int, string) {
	tx, err := db.InitTest().Begin()
	c.Assert(err, IsNil)

	context.Set(middleware.TxKey, tx)
	err = handler(context)
	c.Assert(err, IsNil)

	return tx, rec.Code, rec.Body.String()
}

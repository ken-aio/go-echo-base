package api

import (
	"net/http"

	"github.com/ken-aio/go-echo-base/model"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	. "gopkg.in/check.v1"
)

func (s *ApiSuite) Test_PostXxx(c *C) {
	expectedName := "Test_PostXxx"
	xxx_json := `{"name":"` + expectedName + `"}`
	context, rec := buildContext(c, echo.POST, "/api/v1/xxxs", xxx_json)
	tx, code, resp := requestAPI(c, PostXxx(), context, rec)
	resXxx := new(model.Xxx)
	tx.Select("*").From("xxxs").Where("name = ?", expectedName).LoadStruct(resXxx)
	tx.Rollback()
	logrus.Info(resXxx)

	c.Assert(expectedName, Equals, resXxx.Name)
	c.Assert(http.StatusCreated, Equals, code)
	c.Assert("", Equals, resp)
}

func (s *ApiSuite) Test_GetXxx(c *C) {
	expected := `{"id":1,"name":"name1"}`
	context, rec := buildContext(c, echo.GET, "dummy", "")
	context.SetPath("/api/v1/xxxs/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	tx, code, resp := requestAPI(c, GetXxx(), context, rec)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(expected, Equals, resp)
}

func (s *ApiSuite) Test_GetXxxs(c *C) {
	expected := `[{"id":1,"name":"name1"},{"id":2,"name":"name2"}]`
	context, rec := buildContext(c, echo.GET, "dummy", "")
	context.SetPath("/api/v1/xxxs")
	tx, code, resp := requestAPI(c, GetXxxs(), context, rec)
	tx.Rollback()

	c.Assert(http.StatusOK, Equals, code)
	c.Assert(expected, Equals, resp)
}

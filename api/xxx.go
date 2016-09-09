package api

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/ken-aio/go-echo-base/model"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetXxx() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		xxx := new(model.Xxx)
		if err := xxx.Load(tx, id); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Xxx does not exists. id: " + strconv.FormatInt(id, 10))
		}
		return c.JSON(fasthttp.StatusOK, xxx)
	}
}

func GetXxxs() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)

		position := c.QueryParam("position")
		xxxs := new(model.Xxxs)
		if err = xxxs.Load(tx, position); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "Xxx does not exists.")
		}

		return c.JSON(fasthttp.StatusOK, xxxs)
	}
}

func PostXxx() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		m := new(model.Xxx)
		c.Bind(&m)

		tx := c.Get("Tx").(*dbr.Tx)

		xxx := model.NewXxx(m.Name)

		if err := xxx.Save(tx); err != nil {
			logrus.Error(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.NoContent(fasthttp.StatusCreated)
	}
}

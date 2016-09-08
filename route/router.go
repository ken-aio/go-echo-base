package route

import (
	"github.com/ken-aio/echo-base/api"
	"github.com/ken-aio/echo-base/db"
	"github.com/ken-aio/echo-base/handler"
	myMw "github.com/ken-aio/echo-base/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/xxxs", api.PostXxx())
		v1.GET("/xxxs", api.GetXxxs())
		v1.GET("/xxxs/:id", api.GetXxx())
	}
	return e
}

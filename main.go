package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/ken-aio/go-echo-base/route"
	"github.com/labstack/echo/engine/fasthttp"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Init()
	router.Run(fasthttp.New(":8888"))
}

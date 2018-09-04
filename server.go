package main

import (
	"fmt"
	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uuidcode/echo-test2/controller/bookController"
	"net/http"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := fmt.Sprintf("%d.html", code)

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

	c.Logger().Error(err)
}

func main() {
	templateConfig := echotemplate.TemplateConfig{
		Root:         "template",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	}

	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Renderer = echotemplate.New(templateConfig)
	e.Use(middleware.Logger())
	e.Static("/html", "static/html")
	e.File("/favicon.ico", "static/ico/favicon.ico")

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World")
	})

	e.GET("/book", bookController.FindAllBookHandle)

	e.GET("/book/template", bookController.Template)
	e.GET("/book/innerTemplate", bookController.InnerTemplate)

	e.Start(":17002")
}

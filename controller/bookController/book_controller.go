package bookController

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/uuidcode/echo-test2/dao/bookDao"
	"net/http"
)

func FindAllBookHandle(context echo.Context) error {
	bookList := bookDao.FindAll()
	return context.JSON(http.StatusOK, bookList)
}

func Template(context echo.Context) error {
	bookList := bookDao.FindAll()
	return context.Render(http.StatusOK, "book/index", echo.Map{
		"bookList": bookList,
	})
}

func InnerTemplate(context echo.Context) error {
	var b bytes.Buffer
	e := context.Echo()

	bookList := bookDao.FindAll()

	e.Renderer.Render(&b, "book/test.html", echo.Map{
		"bookList": bookList,
	}, context)

	return context.String(http.StatusOK, b.String())
}

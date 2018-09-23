package echoHelper

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Route struct {
		Method         string
		Path           string
		ControllerFunc echo.HandlerFunc
	}
)

type (
	EchoHelper struct {
		_echo *echo.Echo
	}
)

func New(e *echo.Echo) (echoHelper *EchoHelper) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	echoHelper = &EchoHelper{
		_echo: e,
	}
	return
}

func (this *EchoHelper) Serve() {
	this._echo.Logger.Fatal(this._echo.Start(":1323"))
}

func (this *EchoHelper) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case echo.CONNECT:
			this._echo.CONNECT(route.Path, route.ControllerFunc)
		case echo.DELETE:
			this._echo.DELETE(route.Path, route.ControllerFunc)
		case echo.GET:
			this._echo.GET(route.Path, route.ControllerFunc)
		case echo.HEAD:
			this._echo.HEAD(route.Path, route.ControllerFunc)
		case echo.OPTIONS:
			this._echo.OPTIONS(route.Path, route.ControllerFunc)
		case echo.PATCH:
			this._echo.PATCH(route.Path, route.ControllerFunc)
		case echo.POST:
			this._echo.POST(route.Path, route.ControllerFunc)
		case echo.PUT:
			this._echo.PUT(route.Path, route.ControllerFunc)
		case echo.TRACE:
			this._echo.TRACE(route.Path, route.ControllerFunc)
		default:
			panic("unknown route definition")
		}
	}
}

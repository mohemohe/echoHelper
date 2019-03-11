package echoHelper

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Route struct {
		Method         string
		Path           string
		ControllerFunc echo.HandlerFunc
		MiddleWareFuncs *[]echo.MiddlewareFunc
	}

	EchoHelper struct {
		_echo *echo.Echo
	}

	echoHelperOptions struct {
		middlewareFuncs []echo.MiddlewareFunc
	}

	echoHelperOption func(*echoHelperOptions)
)

func New(e *echo.Echo, option ...echoHelperOption) (echoHelper *EchoHelper) {
	opt := echoHelperOptions{}
	for _, o := range option {
		o(&opt)
	}

	if opt.middlewareFuncs == nil {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
	} else {
		for _, o := range opt.middlewareFuncs {
			e.Use(o)
		}
	}

	echoHelper = &EchoHelper{
		_echo: e,
	}
	return
}

func WithCustomMiddleware(middlewareFuncs []echo.MiddlewareFunc) echoHelperOption {
	return func(opt *echoHelperOptions) {
		opt.middlewareFuncs = middlewareFuncs
	}
}

func (this *EchoHelper) Echo() *echo.Echo {
	return this._echo
}

func (this *EchoHelper) Serve() {
	this._echo.Logger.Fatal(this._echo.Start(":1323"))
}

func (this *EchoHelper) RegisterRoutes(routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case echo.CONNECT:
			if route.MiddleWareFuncs != nil {
				this._echo.CONNECT(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.CONNECT(route.Path, route.ControllerFunc)
			}
		case echo.DELETE:
			if route.MiddleWareFuncs != nil {
				this._echo.DELETE(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.DELETE(route.Path, route.ControllerFunc)
			}
		case echo.GET:
			if route.MiddleWareFuncs != nil {
				this._echo.GET(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.GET(route.Path, route.ControllerFunc)
			}
		case echo.HEAD:
			if route.MiddleWareFuncs != nil {
				this._echo.HEAD(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.HEAD(route.Path, route.ControllerFunc)
			}
		case echo.OPTIONS:
			if route.MiddleWareFuncs != nil {
				this._echo.OPTIONS(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.OPTIONS(route.Path, route.ControllerFunc)
			}
		case echo.PATCH:
			if route.MiddleWareFuncs != nil {
				this._echo.PATCH(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.PATCH(route.Path, route.ControllerFunc)
			}
		case echo.POST:
			if route.MiddleWareFuncs != nil {
				this._echo.POST(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.POST(route.Path, route.ControllerFunc)
			}
		case echo.PUT:
			if route.MiddleWareFuncs != nil {
				this._echo.PUT(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.PUT(route.Path, route.ControllerFunc)
			}
		case echo.TRACE:
			if route.MiddleWareFuncs != nil {
				this._echo.TRACE(route.Path, route.ControllerFunc, *route.MiddleWareFuncs...)
			} else {
				this._echo.TRACE(route.Path, route.ControllerFunc)
			}
		default:
			panic("unknown route definition")
		}
	}
}

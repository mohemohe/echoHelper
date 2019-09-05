package echoHelper

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
	"testing"
)

var eh *EchoHelper

func TestEchoHelper_New(t *testing.T) {
	eh = New(echo.New(), WithCustomMiddleware([]echo.MiddlewareFunc{
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `${time_rfc3339_nano} ${method} ${uri} ${status} ${latency_human}` + "\n",
		}),
	}))
	if eh == nil {
		t.Error("echo helper instance is nil")
	}
	if eh._echo == nil {
		t.Error("internal echo instance is nil")
	}
	eh._echo.HideBanner = true
}

func TestEchoHelper_RegisterRoutes(t *testing.T) {
	dummyHandler1 := func(c echo.Context) error {
		return c.HTML(200, "OK")
	}

	dummyHandler2 := func(c echo.Context) error {
		return c.HTML(404, "Not Found")
	}

	dummyMiddleware1 := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.HTML(500, "Internal Server Error")
		}
	}

	t.Run("check initial route state", func(t *testing.T) {
		if eh._echo.Routes() == nil {
			t.Error("internal echo routes instance is nil")
		}
		if len(eh._echo.Routes()) != 0 {
			t.Error("internal echo routes seems invalid")
		}
	})

	t.Run("set one route", func(t *testing.T) {
		eh.RegisterRoutes([]Route{
			{echo.GET, "/", dummyHandler1, nil},
		})

		if eh._echo.Routes() == nil {
			t.Error("internal echo routes instance is nil")
		}
		if len(eh._echo.Routes()) != 1 {
			t.Error("internal echo routes seems invalid")
		}
	})

	t.Run("append two routes", func(t *testing.T) {
		eh.RegisterRoutes([]Route{
			{echo.GET, "/1", dummyHandler1, nil},
			{echo.GET, "/2", dummyHandler1, nil},
		})

		if eh._echo.Routes() == nil {
			t.Error("internal echo routes instance is nil")
		}
		if len(eh._echo.Routes()) != 3 {
			t.Error("internal echo routes seems invalid")
		}
	})


	t.Run("update two routes", func(t *testing.T) {
		eh.RegisterRoutes([]Route{
			{echo.GET, "/1", dummyHandler2, nil},
			{echo.GET, "/2", dummyHandler2, &[]echo.MiddlewareFunc{dummyMiddleware1}},
		})

		if eh._echo.Routes() == nil {
			t.Error("internal echo routes instance is nil")
		}
		if len(eh._echo.Routes()) != 3 {
			t.Error("internal echo routes seems invalid")
		}
	})
}

func TestEchoHelper_Serve(t *testing.T) {
	t.Run("serve without address", func(t *testing.T) {
		go eh.Serve()
		defer eh.Shutdown()

		url := "http://localhost:1323"

		t.Run("GET /", func(t *testing.T) {
			resp, _ := http.Get(url + "/")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "OK" {
				t.Error("internal echo routes seems invalid")
			}
		})
		t.Run("GET /1", func(t *testing.T) {
			resp, _ := http.Get(url + "/1")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "Not Found" {
				t.Error("internal echo routes seems invalid")
			}
		})
		t.Run("GET /2", func(t *testing.T) {
			resp, _ := http.Get(url + "/2")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "Internal Server Error" {
				t.Error("internal echo routes seems invalid")
			}
		})
	})

	t.Run("serve with address", func(t *testing.T) {
		TestEchoHelper_New(t)
		TestEchoHelper_RegisterRoutes(t)
		go eh.Serve(":1324")

		url := "http://localhost:1324"

		t.Run("GET /", func(t *testing.T) {
			resp, _ := http.Get(url + "/")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "OK" {
				t.Error("internal echo routes seems invalid")
			}
		})
		t.Run("GET /1", func(t *testing.T) {
			resp, _ := http.Get(url + "/1")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "Not Found" {
				t.Error("internal echo routes seems invalid")
			}
		})
		t.Run("GET /2", func(t *testing.T) {
			resp, _ := http.Get(url + "/2")
			defer resp.Body.Close()

			byteArray, _ := ioutil.ReadAll(resp.Body)
			if string(byteArray) != "Internal Server Error" {
				t.Error("internal echo routes seems invalid")
			}
		})
	})
}
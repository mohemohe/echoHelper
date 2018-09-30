package echoHelper

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"testing"
)

var eh *EchoHelper

func TestEchoHelper_New(t *testing.T) {
	eh = New(echo.New())
	if eh == nil {
		t.Error("echo helper instance is nil")
	}
	if eh._echo == nil {
		t.Error("internal echo instance is nil")
	}
}

func TestEchoHelper_RegisterRoutes(t *testing.T) {
	dummyHandler1 := func(c echo.Context) error {
		return c.HTML(200, "OK")
	}

	dummyHandler2 := func(c echo.Context) error {
		return c.HTML(404, "Not Found")
	}

	eh.RegisterRoutes([]Route{
		{echo.GET, "/", dummyHandler1},
	})

	if eh._echo.Routes() == nil {
		t.Error("internal echo routes instance is nil")
	}
	if len(eh._echo.Routes()) != 1 {
		t.Error("internal echo routes seems invalid")
	}

	eh.RegisterRoutes([]Route{
		{echo.GET, "/1", dummyHandler1},
		{echo.GET, "/2", dummyHandler1},
	})

	if len(eh._echo.Routes()) != 3 {
		t.Error("internal echo routes seems invalid")
	}

	eh.RegisterRoutes([]Route{
		{echo.GET, "/1", dummyHandler2},
		{echo.GET, "/2", dummyHandler2},
	})

	if len(eh._echo.Routes()) != 3 {
		t.Error("internal echo routes seems invalid")
	}
}

func TestEchoHelper_Serve(t *testing.T) {
	go eh.Serve()

	url := "http://localhost:1323"

	func() {
		resp, _ := http.Get(url + "/")
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		if string(byteArray) != "OK" {
			t.Error("internal echo routes seems invalid")
		}
	}()
	func() {
		resp, _ := http.Get(url + "/1")
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		if string(byteArray) != "Not Found" {
			t.Error("internal echo routes seems invalid")
		}
	}()
	func() {
		resp, _ := http.Get(url + "/2")
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		if string(byteArray) != "Not Found" {
			t.Error("internal echo routes seems invalid")
		}
	}()
}
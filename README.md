# echoHelper
routing helper for labstack/echo

## usage

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/echoHelper/v4"
	"controllers"
)

func main() {
	eh := echoHelper.New(echo.New())
	eh.RegisterRoutes([]echoHelper.Route{
		{echo.GET, "/api/v1/auth", auth.Get, nil},
		{echo.POST, "/api/v1/auth", auth.Post, nil},
		{echo.GET, "/api/v1/me", me.Get, nil},
		{echo.PUT, "/api/v1/me", me.Put, nil},
		{echo.GET, "/api/v1/users", users.Find, nil},
		{echo.GET, "/api/v1/users/:userId", users.Get, nil},
	})
	eh.Serve()
}
```
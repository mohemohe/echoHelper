# echoHelper
routing helper for labstack/echo

## usage

```go
package main

import (
	"github.com/labstack/echo"
	"github.com/mohemohe/echoHelper"
	"controllers"
)

func main() {
	eh := echoHelper.New(echo.New())
	eh.RegisterRoutes([]echoHelper.Route{
		{echo.GET, "/api/v1/auth", auth.Get},
		{echo.POST, "/api/v1/auth", auth.Post},
		{echo.GET, "/api/v1/me", me.Get},
		{echo.PUT, "/api/v1/me", me.Put},
		{echo.GET, "/api/v1/users", users.Find},
		{echo.GET, "/api/v1/users/:userId", users.Get},
	})
	eh.Serve()
}
```
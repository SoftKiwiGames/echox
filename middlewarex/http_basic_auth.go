package middlewarex

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func BasicAuth(httpUser string, httpPassword string, realm string) echo.MiddlewareFunc {
	return middleware.BasicAuthWithConfig(
		middleware.BasicAuthConfig{
			Validator: func(username, password string, c echo.Context) (bool, error) {
				if subtle.ConstantTimeCompare([]byte(username), []byte(httpUser)) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(httpPassword)) == 1 {
					return true, nil
				}
				return false, nil
			},
			Realm: realm,
		},
	)
}

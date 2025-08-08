package routes

import (
	"github.com/labstack/echo"
	"github.com/satyakarki/did-api/controllers"
)

type Routes struct {
	echo    *echo.Echo
	didCtrl *controllers.DIDController
}

func New(e *echo.Echo, didCtrl *controllers.DIDController) *Routes {
	return &Routes{
		echo:    e,
		didCtrl: didCtrl,
	}
}

func (r *Routes) Init() {
	e := r.echo

	// Define your routes here
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	g := e.Group("/v1/api/did")
	g.POST("/create", r.didCtrl.CreateDID)

	g.GET("/resolve/:id", r.didCtrl.ResolveDID)          // Alias for resolve
	g.POST("/add-key/:id", r.didCtrl.AddKeyToDIDHandler) // you have a method to add keys to DID
	g.POST("/sign", r.didCtrl.SignMessageHandler)
	g.POST("/verify", r.didCtrl.VerifyMessageHandler)
}

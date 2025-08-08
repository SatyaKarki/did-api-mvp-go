package cmd

import (
	"github.com/labstack/echo"
	"github.com/satyakarki/did-api/config"
	"github.com/satyakarki/did-api/controllers"
	"github.com/satyakarki/did-api/routes"
	"github.com/satyakarki/did-api/server"
	"github.com/satyakarki/did-api/services/did"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: Serve,
}

func Serve(cmd *cobra.Command, args []string) {
	//dbClient := conn.Db()

	// services
	didService := did.NewDIDService()

	//get controller
	didCtrl := controllers.NewDIDController(didService)

	// Initialize the server
	echoServer := echo.New()
	server := server.New(echoServer)

	//register routes
	routes := routes.New(echoServer, didCtrl)
	routes.Init()

	// Start the server
	server.Start(config.App().Port)
}

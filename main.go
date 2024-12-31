package main

import (
	"confab/server"
	"confab/utils"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type Message struct {
	message interface{}
}

func customHandler(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight (OPTIONS) request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Respond with JSON
	json.NewEncoder(w).Encode(Message{message: "Hello from the Wails server!"})
}

func main() {
	go server.NewServer()

	//listen to custom server from server package
	// server.NewServer().StartServer()
	// Create an instance of the app structure
	app := NewApp()

	//grab server info to feed to phone
	myIp := utils.NewIpDetailsGrabber()
	fmt.Println("MY IP: ", myIp.Grabber.GetServerInfo())
	fmt.Println("MY IP full Url: ", myIp.Grabber.GetServerInfo().ServerUrl)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "confab",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},

		Logger:   nil,          // Disable Wails' logger entirely
		LogLevel: logger.ERROR, // Show only critical errors
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

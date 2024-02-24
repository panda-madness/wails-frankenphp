package main

import (
	"embed"
	"net/http"

	"github.com/dunglas/frankenphp"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type FrankenHandler struct{}

func (h FrankenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := frankenphp.NewRequestWithContext(
		r,
		frankenphp.WithRequestDocumentRoot("/Users/margulan/Coding/Personal/wails-frankenphp/root/public/index.php", false),
	)

	if err != nil {
		panic(err)
	}

	if err := frankenphp.ServeHTTP(w, req); err != nil {
		panic(err)
	}
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	if err := frankenphp.Init(); err != nil {
		panic(err)
	}

	defer frankenphp.Shutdown()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-frankenphp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				return FrankenHandler{}
			},
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

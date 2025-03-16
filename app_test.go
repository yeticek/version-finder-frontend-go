package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"Year":  time.Now().Year(),
		})
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		// Mock data for testing
		apps := []AppInfo{
			{Source: "GitHub", AppName: "App1", Version: "1.0.0", LatestVersion: "1.0.1"},
		}

		// Render the data in an HTML table
		return c.Render("table", fiber.Map{
			"Title": "App Versions",
			"Apps":  apps,
			"Year":  time.Now().Year(),
		})
	})

	return app
}

func TestIndexRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestJsonRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

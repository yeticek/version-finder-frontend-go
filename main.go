package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type AppInfo struct {
	Source        string `json:"source"`
	AppName       string `json:"app_name"`
	Version       string `json:"version"`
	LatestVersion string `json:"latest_version"`
}

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Make an HTTP GET request to the API endpoint
		resp, err := http.Get("http://localhost:9999/api/response")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching data")
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(resp.Body)

		// Parse the JSON response
		var apps []AppInfo
		if err := json.NewDecoder(resp.Body).Decode(&apps); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error parsing JSON")
		}

		// Render the data in an HTML table
		return c.Render("table", fiber.Map{
			"Title": "App Versions",
			"Apps":  apps,
			"Year":  time.Now().Year(),
		})
	})

	log.Fatal(app.Listen(":9998"))
}

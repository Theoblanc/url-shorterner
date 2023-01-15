package routes

import (
	"github.com/Theoblanc/url-shortener/shortener"
	"github.com/gofiber/fiber/v2"
)

var routeName string = "/shortens"

// ShortenRouter is shorten router
type ShortenRouter struct {
	app        fiber.Router
	Controller shortener.Controller
}

// ShortenRouters group shortenRouters
func (sr *ShortenRouter) ShortenRouters() {
	sr.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	sr.app.Get(routeName, sr.Controller.Create)
	sr.app.Post(routeName, sr.Controller.GetAll)
	sr.app.Post(routeName+":url", sr.Controller.GetByUrl)
}

package shortener

import (
	"time"

	"github.com/Theoblanc/url-shortener/helpers"
	"github.com/gofiber/fiber/v2"
)

// Controller shortener controller struct
type Controller struct {
	shortenerservice Service
}

// CreateShortenDTO create shorten DTO
type CreateShortenDTO struct {
	url         string
	customShort string
	expiry      time.Time
}

// Create function call Create shortener service
func (c *Controller) Create(ctx *fiber.Ctx) (string, error) {
	dto := new(CreateShortenDTO)

	if err := ctx.BodyParser(dto); err != nil {
		return "", ctx.SendString(err.Error())
	}

	enforceURL := helpers.EnforceHTTP(dto.url)
	dto.customShort = helpers.RemoveDomainError(enforceURL)
	dto.expiry = time.Now().Local().AddDate(1, 0, 0)

	shortURL, err := c.shortenerservice.Create(dto)
	if err != nil {
		return "", ctx.SendString(err.Error())
	}

	return shortURL, nil
}

// GetShorteerDTO get shortener DTO
type GetShorteerDTO struct {
	ID  string
	URL string
}

// GetByURL get shortener by url at Controller
func (c *Controller) GetByURL(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	shoterner, err := c.shortenerservice.GetByURL(url)
	if err != nil {
		ctx.Status(400)
	}

	return ctx.JSON(shoterner)
}

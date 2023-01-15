package shortener

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// ShortenerController controller interface
// type ShortenerController interface {
// 	Create(ctx *fiber.Ctx) (string, error)
// 	GetAll(ctx *fiber.Ctx)
// 	GetByURL(ctx *fiber.Ctx)
// }

// Controller shortener controller struct
type Controller struct {
	shortenerservice ShortenerService
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

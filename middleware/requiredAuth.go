package middleware

import "github.com/gofiber/fiber/v2"

func RequiredAuth(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	c.Set("X-XSS-Protection", "1; mode-block")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-DNS-Prefetch-Control", "off")

	// Go to next middleware:
	return c.Next()
}

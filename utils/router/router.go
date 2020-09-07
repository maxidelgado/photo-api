package router

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github/maxidelgado/photo-api/utils/ctxhelper"
	"github/maxidelgado/photo-api/utils/logger"
	"go.uber.org/zap"
)

const (
	ApiKeyHeader = "X-Api-Key"
	UserIdHeader = "X-User-Id"

	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultIdleTimeout  = 10 * time.Second
)

type router struct {
	app    *fiber.App
	logger *zap.Logger
}

func New() *router {
	app := fiber.New(&fiber.Settings{
		ErrorHandler: defaultErrorHandler,
	})

	r := &router{
		app:    app,
		logger: logger.Logger(&logger.Config{Level: "info"}),
	}

	// Show Fiber logo on console for debug mode
	app.Settings.DisableStartupMessage = true

	app.Use(
		middleware.RequestID(),
		setupContext,
		logRequest,
		logResponse,
	)

	return r
}

func (r *router) Engine() *fiber.App {
	return r.app
}

func setupContext(c *fiber.Ctx) {
	rid := c.Fasthttp.Response.Header.Peek(fiber.HeaderXRequestID)

	ch := ctxhelper.WithContext(c.Context())
	ch.SetRequestId(string(rid))

	c.Locals(ctxhelper.Key, ch)
	c.Next()
}

func logRequest(c *fiber.Ctx) {
	log := logger.WithContext(c.Context())
	log.Info(
		"request",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
	)
	c.Next()
}

func logResponse(c *fiber.Ctx) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	log := logger.WithContext(c.Context())
	log.Info("response",
		zap.Int64("rt", duration.Milliseconds()),
		zap.Int("status", c.Fasthttp.Response.StatusCode()),
		zap.String("body", string(c.Fasthttp.Response.Body())),
	)
}

func defaultErrorHandler(ctx *fiber.Ctx, err error) {
	code := http.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	ctx.Status(code).SendString(err.Error())
}

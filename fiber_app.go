package fiber_base_app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wI2L/jettison"
)

type Config struct {
	FiberConfig      fiber.Config
	CompressionLevel CompressionLevel
}

var configDefault = Config{
	FiberConfig: fiber.Config{
		AppName:     "",
		JSONEncoder: CustomJSONEncoder,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(CustomHTTPError); ok {
				return ctx.Status(e.Code).JSON(e)
			}
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			ce := CustomHTTPError{
				Code:    code,
				Message: err.Error(),
			}
			err = ctx.Status(code).JSON(ce)
			if err != nil {
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}
			return nil
		},
	},
	CompressionLevel: LevelDisabled,
}

// get config default values
func getConfig(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return configDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.CompressionLevel < LevelDisabled || cfg.CompressionLevel > LevelBestCompression {
		cfg.CompressionLevel = configDefault.CompressionLevel
	}

	if cfg.FiberConfig.AppName == "" {
		cfg.FiberConfig.AppName = configDefault.FiberConfig.AppName
	}

	if cfg.FiberConfig.JSONEncoder == nil {
		cfg.FiberConfig.JSONEncoder = configDefault.FiberConfig.JSONEncoder
	}

	if cfg.FiberConfig.ErrorHandler == nil {
		cfg.FiberConfig.ErrorHandler = configDefault.FiberConfig.ErrorHandler
	}

	return cfg
}

func NewFiberApp(appOptions Config) *fiber.App {
	config := getConfig(appOptions)
	f := fiber.New(config.FiberConfig)

	if config.CompressionLevel != LevelDisabled {
		f.Use(compress.New(compress.Config{
			Level: config.CompressionLevel.GetFiberCompressionLevel(),
		}))
	}

	// favicon
	f.Use(favicon.New())

	f.Use(recover.New(recover.Config{EnableStackTrace: true}))

	return f
}

func CustomJSONEncoder(v any) ([]byte, error) {
	//return sonic.Marshal(v)
	return jettison.MarshalOpts(v, jettison.UnsortedMap())
}

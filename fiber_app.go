package fiber_base_app

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wI2L/jettison"
)

type Config struct {
	Name             string           `json:"name"`
	Version          string           `json:"version"`
	FiberConfig      fiber.Config     `json:"-"`
	CompressionLevel CompressionLevel `json:"-"`
}

var configDefault = Config{
	CompressionLevel: LevelDisabled,
	Version:          "unknown",
	FiberConfig: fiber.Config{
		AppName:               "",
		JSONEncoder:           CustomJSONEncoder,
		ProxyHeader:           "X-Forwarded-For",
		DisableStartupMessage: true,
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

	if cfg.Version == "" {
		cfg.Version = configDefault.Version
	}

	if cfg.FiberConfig.AppName == "" {
		if cfg.Name != "" {
			cfg.FiberConfig.AppName = cfg.Name
		} else {
			cfg.FiberConfig.AppName = configDefault.FiberConfig.AppName
		}
	}

	if cfg.FiberConfig.JSONEncoder == nil {
		cfg.FiberConfig.JSONEncoder = configDefault.FiberConfig.JSONEncoder
	}

	if cfg.FiberConfig.ErrorHandler == nil {
		cfg.FiberConfig.ErrorHandler = configDefault.FiberConfig.ErrorHandler
	}

	if cfg.FiberConfig.ProxyHeader == "" {
		cfg.FiberConfig.ProxyHeader = configDefault.FiberConfig.ProxyHeader
	}

	return cfg
}

type App struct {
	Name     string    `json:"name"`
	Version  string    `json:"version"`
	Started  time.Time `json:"started"`
	Uptime   float64   `json:"uptime"`
	fiberApp *fiber.App
}

var appLock = &sync.Mutex{}
var appInstance *App

func NewFiberApp(appOptions Config) *fiber.App {
	if appInstance == nil {
		appLock.Lock()
		defer appLock.Unlock()
		if appInstance == nil {
			config := getConfig(appOptions)
			f := fiber.New(config.FiberConfig)

			if config.CompressionLevel != LevelDisabled {
				f.Use(compress.New(compress.Config{
					Level: config.CompressionLevel.GetFiberCompressionLevel(),
				}))
			}

			app := App{
				Name:     config.Name,
				Version:  config.Version,
				Started:  time.Now(),
				fiberApp: f,
			}

			// favicon
			f.Use(favicon.New())
			// recover
			f.Use(recover.New(recover.Config{EnableStackTrace: true}))

			f.Get("/", func(c *fiber.Ctx) error {
				app.Uptime = GetUptime()
				return c.JSON(app)
			})

			appInstance = &app
		}
	}

	return appInstance.fiberApp
}

func GetUptime() float64 {
	now := time.Now()
	return now.Sub(appInstance.Started).Seconds()
}

func GetFiber() *fiber.App {
	return appInstance.fiberApp
}

func CustomJSONEncoder(v any) ([]byte, error) {
	//return sonic.Marshal(v)
	return jettison.MarshalOpts(v, jettison.UnsortedMap())
}

# Fiber Base App

Good Defaults for a fiber rest API

## Installation

```bash
go get github.com/gofiber/fiber/v2
go get dev.azure.com/xompass/xompass-cloud/_git/fiber_base_app.git
```

## Usage

```go
package main

import (
    "dev.azure.com/xompass/xompass-cloud/_git/fiber_base_app.git"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

func main() {
    app := fiber_base_app.NewFiberApp(fiber_base_app.Config{
        // Set your fiber config here
        FiberConfig: fiber.Config{
            AppName: "License plates microservice",
        },
        // Optional: define a compression level
        CompressionLevel: fiber_base_app.LevelDefault,
    })
    api := app.Group("/api")
    
    // Add your routes here
    api.Get("/hello", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    
    port := 8000
    log.Fatal(app.Listen(fmt.Sprint(":", port)))
}
```
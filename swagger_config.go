package fiber_base_app

import (
	"github.com/gofiber/swagger"
	"github.com/xompass/fiber_base_app/swagger_defaults"
)

func GetDefaultSwaggerConfig() swagger.Config {

	return swagger.Config{
		SyntaxHighlight: &swagger.SyntaxHighlightConfig{Activate: false},
		CustomStyle:     swagger_defaults.DefaultTheme,
	}
}

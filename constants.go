package fiber_base_app

import "github.com/gofiber/fiber/v2/middleware/compress"

type CompressionLevel int

// Represents compression level that will be used in the middleware
const (
	LevelDisabled        CompressionLevel = -1
	LevelDefault         CompressionLevel = 0
	LevelBestSpeed       CompressionLevel = 1
	LevelBestCompression CompressionLevel = 2
)

func (receiver CompressionLevel) GetFiberCompressionLevel() compress.Level {
	switch receiver {
	case LevelDefault:
		return compress.LevelDefault
	case LevelBestSpeed:
		return compress.LevelBestSpeed
	case LevelBestCompression:
		return compress.LevelBestCompression
	default:
		return compress.LevelDisabled
	}
}

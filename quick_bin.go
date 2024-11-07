package QuickBin

import (
	"QuickBin/internal/backend"
	"QuickBin/internal/config"
	"QuickBin/internal/keys"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	cfg := config.LoadOrSaveDefault()
	store, err := backend.NewStorage(cfg)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Setup Rate-limit middleware - https://github.com/JGLTechnologies/gin-rate-limit
	rateLimitStore := ratelimit.InMemoryStore(
		&ratelimit.InMemoryOptions{
			Rate:  cfg.RateLimitConfig.Duration,
			Limit: cfg.RateLimitConfig.Max,
		})

	mw := ratelimit.RateLimiter(rateLimitStore, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// Setup Routes
	// POST / - Saves data, rate-limited
	router.POST("/", mw, func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.String(500, "Error reading data")
			return
		}

		if len(data) > cfg.MaxUploadSize {
			c.String(400, "Data too large")
			return
		}

		key := keys.GenerateRandomKey(cfg.KeyLength, cfg.KeyNamespace)
		err = store.Save(key, data)
		if err != nil {
			c.String(500, "Error saving data")
			return
		}

		fmt.Printf("Request to save data for key: %s\n", key)

		c.String(200, key)
	})

	// GET /:key - Gets data for key
	router.GET("/:key", func(c *gin.Context) {
		c.String(200, "Not implemented")
	})

	// GET /raw/:key - Gets raw data for key
	router.GET("/raw/:key", func(c *gin.Context) {
		key := c.Param("key")
		data, err := store.Get(key)
		if err != nil {
			c.String(500, "Error getting data")
			return
		}
		fmt.Printf("Request to get data for key: %s\n", key)
		c.Data(200, "text/plain", data)
	})

	fmt.Printf("Starting server on %s:%d\n", cfg.Host, cfg.Port)
	err = router.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

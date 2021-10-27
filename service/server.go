package service

import (
	"crypto/tls"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	log "go.vocdoni.io/dvote/log"
)

// Run starts a server and listens for requests
func Run() {
	configFile := viper.GetString("config")
	port := viper.GetInt("port")
	useTLS := viper.GetBool("tls")
	cert := viper.GetString("cert")
	key := viper.GetString("key")

	// info
	if useTLS {
		if cert == "" || key == "" {
			log.Fatal("The certificate and key file are needed to run with TLS enabled")
		}
		log.Info("TLS enabled")
	}

	// info
	if configFile != "" {
		log.Infow("Using config file", configFile)
	}

	// Service set up
	app := fiber.New()

	defineEndpoints(app)

	if useTLS {
		// Read TLS certificate
		cer, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
		addr := fmt.Sprintf(":%d", port)

		// Create custom listener
		ln, err := tls.Listen("tcp", addr, tlsConfig)
		if err != nil {
			log.Fatal(err)
		}

		log.Warnw("Listening TLS on", addr)
		log.Fatal(app.Listener(ln))
	} else {
		addr := fmt.Sprintf(":%d", port)
		log.Warnw("Listening HTTP on", addr)
		log.Fatal(app.Listen(addr))
	}
}

func defineEndpoints(app *fiber.App) {
	// CORS
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		return ctx.Next()
	})

	app.Options("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		return ctx.SendStatus(fiber.StatusOK)
	})

	// API endpoints
	app.Get("/v1/:param", handleV1Get)
}

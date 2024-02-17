package main

import (
	"HoneyHollow/job"
	"HoneyHollow/server/config"
	"HoneyHollow/server/routers"
	"embed"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
)

//go:embed frontend/dist/*
var webDirStatic embed.FS

//go:embed frontend/dist/index.html
var indexStatic embed.FS

func main() {
	config.InitConfig()
	prefork := true
	if config.RunMode == "dev" {
		prefork = false
	}
	app := fiber.New(fiber.Config{
		Prefork:       prefork,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "StupidBear Studio",
		AppName:       "Honey Hollow Website",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == 404 || code == 403 {
				return ctx.Redirect("/")
			}
			return ctx.Status(code).SendString(e.Error())
		},
	})
	app.Use(recover2.New())

	if config.RunMode == "dev" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	app.Get("/", filesystem.New(filesystem.Config{
		Root:       http.FS(indexStatic),
		PathPrefix: "frontend/dist",
	}))

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(webDirStatic),
		PathPrefix: "frontend/dist/assets",
	}))

	routers.RegisterApiRouters(app)

	go job.ProcessBookmarks()

	err := app.Listen(config.ListenAddress)
	if err != nil {
		log.Info(err.Error())
		return
	}
}

package routers

import (
	"HoneyHollow/server/handles"

	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouters(app *fiber.App) {
	router := app.Group("/api")
	//router.Get("/", handles.Index)
	router.Get("/bookmarks", handles.GetAllBookmarks)
	router.Post("/bookmarks", handles.CreateBookmarks)
	router.Put("/view/:bid", handles.View)
	router.Put("/unView/:bid", handles.UnView)
	router.Delete("/bookmarks/:bid", handles.DeleteBookmarks)
	router.Get("/bookmarks/count", handles.CountUnViewBookmarks)
}

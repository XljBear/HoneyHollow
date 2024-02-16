package routers

import (
	"HoneyHollow/server/handles"

	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouters(app *fiber.App) {
	router := app.Group("/api")
	router.Get("/bookmarks", handles.GetAllBookmarks)
	router.Get("/bookmarks/count", handles.CountUnViewBookmarks)
	router.Get("/bookmarks/:bid", handles.GetBookMarks)
	router.Post("/bookmarks", handles.CreateBookmarks)
	router.Put("/bookmarks", handles.UpdateBookmarks)
	router.Delete("/bookmarks/:bid", handles.DeleteBookmarks)

	router.Put("/view/:bid", handles.View)
	router.Put("/unView/:bid", handles.UnView)
}

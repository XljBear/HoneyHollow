package handles

import (
	"HoneyHollow/server/common"
	"HoneyHollow/server/consts"
	"HoneyHollow/server/models"
	"HoneyHollow/server/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBookmarks(c *fiber.Ctx) error {
	type requestStruct struct {
		Url    string `json:"url"`
		Title  string `json:"title"`
		Remark string `json:"remark"`
		Sort   int    `json:"sort"`
	}
	req := &requestStruct{}
	err := c.BodyParser(req)
	if err != nil {
		return c.JSON(common.MakeResponse(consts.ResponseRequestInvalid, "", nil))
	}
	if req.Url == "" || !common.IsValidURL(req.Url) {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, consts.MsgCreateBookmarksUrlInvalid, nil))
	}
	bookmarks := &models.Bookmarks{
		Url:    req.Url,
		Title:  req.Title,
		Remark: req.Remark,
		Sort:   req.Sort,
	}
	result, err := services.CreateBookmarks(bookmarks)
	if err != nil {
		log.Info(err)
		return c.JSON(common.MakeResponse(consts.ResponseStatusSystemError, "", nil))
	}
	if result == nil {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, "", nil))
	}
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", nil))
}

func GetAllBookmarks(c *fiber.Ctx) error {
	viewed := c.Query("v", "false")
	allBookmarks := services.GetAllBookmarks(viewed == "true")
	var allBookmarksWebJson []models.BookmarksJson
	allBookmarksWebJson = make([]models.BookmarksJson, 0)
	for _, b := range allBookmarks {
		allBookmarksWebJson = append(allBookmarksWebJson, b.ToWebJsonStruct())
	}
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", allBookmarksWebJson))
}

func View(c *fiber.Ctx) error {
	bookmarksIdStr := c.Params("bid")
	bookmarksId, err := primitive.ObjectIDFromHex(bookmarksIdStr)
	if err != nil {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, consts.MsgBookmarksIdInvalid, nil))
	}
	if !services.ViewBookmarks(bookmarksId) {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, "", nil))
	}
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", nil))
}

func UnView(c *fiber.Ctx) error {
	bookmarksIdStr := c.Params("bid")
	bookmarksId, err := primitive.ObjectIDFromHex(bookmarksIdStr)
	if err != nil {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, consts.MsgBookmarksIdInvalid, nil))
	}
	if !services.UnViewBookmarks(bookmarksId) {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, "", nil))
	}
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", nil))
}

func DeleteBookmarks(c *fiber.Ctx) error {
	bookmarksIdStr := c.Params("bid")
	bookmarksId, err := primitive.ObjectIDFromHex(bookmarksIdStr)
	if err != nil {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, consts.MsgBookmarksIdInvalid, nil))
	}
	if !services.DeleteBookmarks(bookmarksId) {
		return c.JSON(common.MakeResponse(consts.ResponseStatusFailed, "", nil))
	}
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", nil))
}

func CountUnViewBookmarks(c *fiber.Ctx) error {
	count := services.CountUnViewBookmarks()
	return c.JSON(common.MakeResponse(consts.ResponseStatusOk, "", count))
}

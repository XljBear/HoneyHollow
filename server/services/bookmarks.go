package services

import (
	"HoneyHollow/server/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindBookmarksById(id primitive.ObjectID) *models.Bookmarks {
	bookmarks := models.Bookmarks{}
	c := models.BookmarksCollection.GetCollection()
	err := c.Find(context.Background(), bson.M{"_id": id}).One(&bookmarks)
	if err != nil {
		return nil
	}
	return &bookmarks
}

func CreateBookmarks(bookmarks *models.Bookmarks) (*models.Bookmarks, error) {
	c := models.BookmarksCollection.GetCollection()
	r, err := c.InsertOne(context.Background(), bookmarks)
	if err != nil {
		return nil, err
	}
	return FindBookmarksById(r.InsertedID.(primitive.ObjectID)), nil
}

func GetAllBookmarks(isViewed bool) []models.Bookmarks {
	var resultBookmarks []models.Bookmarks
	resultBookmarks = make([]models.Bookmarks, 0)
	c := models.BookmarksCollection.GetCollection()
	searchOption := bson.M{"viewed": isViewed}
	_ = c.Find(context.Background(), searchOption).Sort("sort", "-updateAt").All(&resultBookmarks)
	return resultBookmarks
}

func UpdateBookmarks(bookmarks *models.Bookmarks) (err error) {
	c := models.BookmarksCollection.GetCollection()
	update := bson.D{{"$set", bookmarks}}
	err = c.UpdateId(context.Background(), bookmarks.Id, update)
	return
}

func ViewBookmarks(id primitive.ObjectID) bool {
	bookmarks := FindBookmarksById(id)
	if bookmarks == nil {
		return false
	}
	if bookmarks.Viewed == true {
		return true
	}
	now := time.Now().UTC()
	bookmarks.ViewTime = &now
	bookmarks.Viewed = true
	err := UpdateBookmarks(bookmarks)
	if err != nil {
		log.Info(err)
		return false
	}
	return true
}

func UnViewBookmarks(id primitive.ObjectID) bool {
	bookmarks := FindBookmarksById(id)
	if bookmarks == nil {
		return false
	}
	if bookmarks.Viewed == false {
		return true
	}
	bookmarks.ViewTime = nil
	bookmarks.Viewed = false
	err := UpdateBookmarks(bookmarks)
	if err != nil {
		log.Info(err)
		return false
	}
	return true
}

func DeleteBookmarks(id primitive.ObjectID) bool {
	bookmarks := FindBookmarksById(id)
	if bookmarks == nil {
		return false
	}
	c := models.BookmarksCollection.GetCollection()
	err := c.RemoveId(context.Background(), bookmarks.Id)
	if err != nil {
		log.Info(err)
		return false
	}
	return true
}

func CountBookmarks(filter bson.M) int64 {
	c := models.BookmarksCollection.GetCollection()
	count, err := c.Find(context.Background(), filter).Count()
	if err != nil {
		log.Info(err)
		return 0
	}
	return count
}

func CountUnViewBookmarks() int64 {
	filter := bson.M{"viewed": false}
	return CountBookmarks(filter)
}

func GetAllNeedProcessBookmarks() []models.Bookmarks {
	var bookmarks []models.Bookmarks
	bookmarks = make([]models.Bookmarks, 0)
	c := models.BookmarksCollection.GetCollection()
	filter := bson.M{"needProcess": true}
	err := c.Find(context.Background(), filter).All(&bookmarks)
	if err != nil {
		log.Info(err)
		return nil
	}
	return bookmarks
}

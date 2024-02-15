package models

import (
	"context"
	"github.com/qiniu/qmgo/field"
	"time"
)

var BookmarksCollection = &CollectionBase{
	CollectionName: "bookmarks",
}

type Bookmarks struct {
	field.DefaultField `bson:",inline"`
	//OwnerId            primitive.ObjectID `bson:"ownerId"`
	Url         string     `bson:"url"`
	Icon        string     `bson:"icon"`
	Title       string     `bson:"title"`
	Remark      string     `bson:"remark"`
	Viewed      bool       `bson:"viewed"`
	ViewTime    *time.Time `bson:"viewTime"`
	Sort        int        `bson:"sort"`
	NeedProcess bool       `bson:"NeedProcess"`
}

type BookmarksJson struct {
	Id        string     `json:"id"`
	InProcess bool       `json:"inProcess"`
	Url       string     `json:"url"`
	Icon      string     `json:"icon"`
	Title     string     `json:"title"`
	Remark    string     `json:"remark"`
	ViewTime  *time.Time `json:"viewTime"`
	Sort      int        `json:"sort"`
	CreateAt  time.Time  `json:"createAt"`
}

func (b *Bookmarks) ToWebJsonStruct() BookmarksJson {
	bookmarksJson := BookmarksJson{
		Id:        b.Id.Hex(),
		InProcess: b.NeedProcess,
		Url:       b.Url,
		Icon:      b.Icon,
		Title:     b.Title,
		Remark:    b.Remark,
		ViewTime:  b.ViewTime,
		Sort:      b.Sort,
		CreateAt:  b.CreateAt,
	}
	return bookmarksJson
}

func (b *Bookmarks) BeforeInsert(ctx context.Context) error {
	b.NeedProcess = true
	return nil
}

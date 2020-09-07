package worker

import (
	"fmt"
	"github/maxidelgado/photo-api/datastore"
	"github/maxidelgado/photo-api/utils/logger"
	"go.uber.org/zap"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	baseUrl      = "http://interview.agileengine.com"
	defaultTTL   = 5 * time.Minute
	pollingDelay = 5 * time.Minute
)

var (
	newHttpClient = resty.New
	newDS         = datastore.New
	log           = logger.Logger(&logger.Config{Level: "info"})
)

func StartPolling() {
	pageNumber := 1
	for range time.Tick(5 * time.Second) {
		hasMore := pollPictures(pageNumber)
		for hasMore {
			pageNumber++
			hasMore = pollPictures(pageNumber)
		}
		pageNumber = 0
	}
}

func pollPictures(pageNumber int) bool {
	result, err := getPaginated(pageNumber)
	if err != nil {
		log.Error("getPaginated()", zap.Error(err))
		return false
	}
	log.Info("getPaginated()", zap.Int("pageNumber", pageNumber))
	ds := newDS()

	for _, picture := range result.Pictures {
		picture, err = getById(picture.ID)
		if err != nil {
			log.Error("getById()", zap.Error(err))
			continue
		}
		log.Info("getById()", zap.Int("pageNumber", pageNumber), zap.String("pic_id", picture.ID))
		ds.Set(picture.ID, picture, defaultTTL)

		putIndicesFromTags(picture.Tags, picture, defaultTTL)
		putIndex(fmt.Sprintf(AuthorIndexPattern, picture.Author), picture, defaultTTL)
		putIndex(fmt.Sprintf(CameraIndexPattern, picture.Camera), picture, defaultTTL)
	}

	return result.HasMore
}

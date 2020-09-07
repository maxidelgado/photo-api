package worker

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	AuthorIndexPattern = "author-%s"
	CameraIndexPattern = "camera-%s"
	TagIndexPattern    = "tag-%s"
)

func putIndex(key string, picture Picture, ttl time.Duration) {
	key = normalizeKey(key)
	ds := newDS()
	content, found := ds.Get(key)
	if !found {
		ds.Set(key, []Picture{picture}, ttl)
		log.Info("putIndex()", zap.String("index", key), zap.String("pic_id", picture.ID))
		return
	}

	pictures, ok := content.([]Picture)
	if !ok {
		return
	}
	pictures = append(pictures, picture)
	ds.Set(key, pictures, ttl)
	log.Info("putIndex()", zap.String("index", key), zap.String("pic_id", picture.ID))
}

func normalizeKey(key string) string {
	key = strings.ToLower(key)
	key = strings.Replace(key, " ", "", -1)
	return key
}

func putIndicesFromTags(tags string, picture Picture, ttl time.Duration) {
	noBlankSpace := strings.Replace(tags, " ", "", -1)
	result := strings.Split(noBlankSpace, "#")
	for _, tag := range result {
		if tag == "" {
			continue
		}
		putIndex(fmt.Sprintf(TagIndexPattern, tag), picture, ttl)
	}
}

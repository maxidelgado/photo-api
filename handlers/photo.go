package handlers

import (
	"github.com/gofiber/fiber"
	"github/maxidelgado/photo-api/datastore"
	"github/maxidelgado/photo-api/utils/logger"
	"net/http"
)

var (
	newDS = datastore.New
	log   = logger.Logger(&logger.Config{Level: "info"})
)

func Search(c *fiber.Ctx) {
	searchTerm := c.Params("searchTerm", "empty")
	if searchTerm == "empty" {
		c.Next(fiber.NewError(http.StatusBadRequest, "searchTerm is required"))
		return
	}

	ds := newDS()
	result, _ := ds.Get(searchTerm)
	c.JSON(result)
}

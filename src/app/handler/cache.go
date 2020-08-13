package handler

import (
	"net/http"
	"search/src/app/cache"
	"search/src/app/common"
)

var imMemory = cache.NewCache()

func (h *Handler) ClearCache(w http.ResponseWriter, _ *http.Request) {
	log.Printf("clear cache request received")
	imMemory.Clear()
	common.RespondJSON(w, http.StatusOK, common.Response{Message: "cache was cleared successfully"})
}
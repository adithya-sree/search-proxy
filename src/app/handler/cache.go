package handler

import (
	"fmt"
	"github.com/gorilla/mux"
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

func (h *Handler) ClearCacheKey(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	key := p["key"]
	log.Printf("clear cache key request received for key [%s]", key)
	imMemory.ClearKey(key)
	common.RespondJSON(w, http.StatusOK, common.Response{Message: fmt.Sprintf("cache for key [%s] was cleared successfully", key)})
}

func (h *Handler) GetCacheKey(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	key := p["key"]
	log.Printf("get cache key request received for key [%s]", key)
	value, err := imMemory.Find(key)
	if err != nil {
		log.Printf("no cache for given key [%s]", key)
		common.RespondError(w, http.StatusNotFound, fmt.Sprintf("cannot find cache for key [%s]", key))
		return
	}
	log.Printf("successfully found cache for key [%s]", key)
	common.RespondJSON(w, http.StatusOK, value)
}
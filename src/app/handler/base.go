package handler

import (
	"net/http"
	"search/src/app/common"
	"time"
)

var startTime time.Time

type UptimeResponse struct {
	StartedTime string        `json:"start-time"`
	Uptime      time.Duration `json:"uptime"`
}

func init() {
	startTime = time.Now()
}

func (h *Handler) Ecv(w http.ResponseWriter, _ *http.Request) {
	log.Println("ecv check received")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Base(w http.ResponseWriter, _ *http.Request) {
	log.Println("base request received")
	common.RespondJSON(w, http.StatusOK, common.Response{Message: "search server is running"})
}

func (h *Handler) Running(w http.ResponseWriter, _ *http.Request) {
	log.Println("running check received")
	common.RespondJSON(w, http.StatusOK, common.Response{Message: "search service running"})
}

func (h *Handler) Uptime(w http.ResponseWriter, _ *http.Request) {
	log.Println("uptime request received")
	common.RespondJSON(w, http.StatusOK, UptimeResponse{
		StartedTime: startTime.Format("2006.01.02 15:04:05"),
		Uptime:      time.Since(startTime),
	})
}

package handler

import (
	"search/src/config"
	"search/src/logger"
)

type Handler struct {
	Config *config.Config
}

var log = logger.GetLogger("handler")

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		Config: c,
	}
}

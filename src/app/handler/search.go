package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"search/src/app/common"
	"search/src/app/service/deezer"
)

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	query := p["query"]
	log.Printf("performing request for query [%s]", query)

	resp, err := imMemory.Find(query)
	if err != nil {
		log.Printf("unable to find query [%s] results in cache, getting from initial data source", query)
		c := deezer.NewClient(*h.Config, query)
		r, err := c.BuildRequest()
		if err != nil {
			log.Printf("error while building request [%v] for query [%s", err, query)
			common.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp, err := c.Execute(r)
		if err != nil {
			log.Printf("error while requesting tracks [%v] for query [%s]", err, query)
			common.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if resp.Data == nil || len(resp.Data) < 1 {
			m := fmt.Sprintf("no content available for given query [%s]", query)
			log.Printf(m)
			common.RespondJSON(w, http.StatusNotFound, common.Response{Message: m})
			return
		}

		imMemory.Add(query, resp)
		log.Printf("successfully executed request for query [%v]", query)
		common.RespondJSON(w, http.StatusOK, resp)
		return
	}

	log.Printf("successfully found query [%s] in cache", query)
	common.RespondJSON(w, http.StatusOK, resp)
}

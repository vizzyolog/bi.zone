package rest

import (
	"context"
	"net/http"

	"app/domain"
	"app/pkg/logger"
	"app/usecases/events"

	"github.com/gorilla/mux"
)

func addEventsAPI(router *mux.Router, ret eventRetriever, logger logger.Logger) {
	ec := &eventController{
		ret: ret,
	}

	router.HandleFunc("/v1/event/{id}", ec.get).Methods(http.MethodGet)
	router.HandleFunc("/v1/events/", ec.search).Methods(http.MethodGet)

}

type eventController struct {
	logger.Logger

	ret eventRetriever
}

func (ec *eventController) search(wr http.ResponseWriter, req *http.Request) {
	posts, err := ec.ret.Search(req.Context(), events.Query{})
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, posts)
}

func (pc *eventController) get(wr http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	post, err := pc.ret.Get(req.Context(), name)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, post)
}

type eventRetriever interface {
	Get(ctx context.Context, name string) (*domain.Event, error)
	Search(ctx context.Context, query events.Query) ([]domain.Event, error)
}

type eventPeristor interface {
	Publish(ctx context.Context, event domain.Event) (*domain.Event, error)
	Delete(ctx context.Context, name string) (*domain.Event, error)
}

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

func (ec *eventController) get(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user, err := ec.ret.Get(req.Context(), vars["name"])
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, user)
}

func (ec *eventController) search(wr http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()["t"]
	users, err := ec.ret.Search(req.Context(), vals, 10)
	if err != nil {
		respondErr(wr, err)
		return
	}

	respond(wr, http.StatusOK, users)
}

type eventRetriever interface {
	Get(ctx context.Context, name string) (*domain.User, error)
	Search(ctx context.Context, tags []string, limit int) ([]domain.User, error)
	VerifySecret(ctx context.Context, name, secret string) bool
}

type postRetriever interface {
	Get(ctx context.Context, name string) (*domain.Event, error)
	Search(ctx context.Context, query events.Query) ([]domain.Event, error)
}

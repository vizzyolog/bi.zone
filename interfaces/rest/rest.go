package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"app/pkg/logger"
	"app/pkg/render"
)

// New initializes the server with routes exposing the given usecases.
func New(logger logger.Logger, reg registration, userRet retriever, eventRet eventRetriever) http.Handler {
	// setup router with default handlers
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)

	// setup api endpoints
	addEventsAPI(router, eventRet, logger)
	addUsersAPI(router, reg, userRet, logger)

	return router
}

func notFoundHandler(wr http.ResponseWriter, req *http.Request) {
	render.JSON(wr, http.StatusNotFound, fmt.Errorf("not found path %v", req.URL.Path))
}

func methodNotAllowedHandler(wr http.ResponseWriter, req *http.Request) {
	render.JSON(wr, http.StatusMethodNotAllowed, fmt.Errorf("not found path %v", req.URL.Path))
}

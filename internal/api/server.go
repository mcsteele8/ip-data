package api

import (
	"encoding/json"
	"errors"
	"ip-data/internal/config"
	"ip-data/internal/spyware"
	"ip-data/tools/werror"
	"ip-data/tools/wlog"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	Spyware *spyware.Spyware
}

func NewServer() *http.Server {

	router := mux.NewRouter()
	svr := &Server{Spyware: spyware.New(&http.Client{})}
	svr.registerHandlers(router)
	handler := cors.Default().Handler(router)
	//Address port should be in a config
	newServer := &http.Server{
		Addr:    config.ServerPort,
		Handler: handler,
	}

	return newServer
}

func (s *Server) registerHandlers(mux *mux.Router) {
	mux.Use(wlog.Middleware)
	mux.HandleFunc("/v1/data/{ip-or-domain}", s.dataHandler).Methods(http.MethodGet)
}

func (s *Server) dataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataRequest := vars["ip-or-domain"]
	if dataRequest == "" {
		werror.DoHttpError(w, http.StatusBadRequest, "Bad Request. Please provide an ip address or domain name")
		return
	}

	info, err := s.Spyware.GetSpywareInfo(dataRequest)
	if err != nil {
		if errors.Is(werror.ErrorBadArgs, err) {
			werror.DoHttpError(w, http.StatusBadRequest, err.Error())
			return
		}

		werror.DoHttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(info); err != nil {
		werror.DoHttpError(w, http.StatusInternalServerError, err.Error())
	}
}

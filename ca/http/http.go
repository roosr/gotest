package dchttp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	TextPlain string = "text/plain; charset=utf-8"
)

type DcHttpServer struct {
	router     *mux.Router
	httpServer *http.Server
}

func New(ipAddress string, port int) *DcHttpServer {

	router := mux.NewRouter()

	httpServer := &http.Server{
		Addr:         ipAddress + ":" + strconv.Itoa(port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	return &DcHttpServer{
		router:     router,
		httpServer: httpServer,
	}
}

func (server *DcHttpServer) AddPostRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
	server.router.HandleFunc(path, handler).Methods("POST")
}

func (server *DcHttpServer) AddGetRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
	server.router.HandleFunc(path, handler).Methods("GET")
}

func GetPathParam(request *http.Request, name string) string {
	params := mux.Vars(request)
	return params[name]
}

func WriteStringResponse(w http.ResponseWriter, text string) {
	WriteStringResponseStatusType(w, text, http.StatusOK, TextPlain)
}

func WriteStringResponseType(w http.ResponseWriter, text string, contentType string) {
	WriteStringResponseStatusType(w, text, http.StatusOK, contentType)
}

func WriteStringResponseStatus(w http.ResponseWriter, text string, statusCode int) {
	WriteStringResponseStatusType(w, text, statusCode, TextPlain)
}

func WriteStringResponseStatusType(w http.ResponseWriter, text string,
	statusCode int, contentType string) {

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	io.WriteString(w, text)
}

func WriteJsonResponse(w http.ResponseWriter, res interface{}) {
	WriteJsonResponseStatus(w, res, http.StatusOK)
}

func WriteJsonResponseStatus(w http.ResponseWriter, res interface{}, statusCode int) {
	jsonData, err := json.Marshal(res)
	if err != nil {
		// TODO(roos): log error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func (server *DcHttpServer) Run() {
	server.httpServer.ListenAndServe()
}

func (server *DcHttpServer) Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.httpServer.Shutdown(ctx)
}

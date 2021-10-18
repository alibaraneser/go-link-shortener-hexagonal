package api

import (
	"errors"
	json2 "github.com/alibaraneser/go-link-shortener-hexagonal/internal/adapters/serializer/json"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/ports"
	"io/ioutil"
	"log"
	"net/http"
)

type RedirectHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	redirectService ports.RedirectService
}

func NewHandler(redirectService ports.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	redirect, err := h.redirectService.Find(code)

	if err != nil {
		if errors.Is(err, domain.RedirectNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	handleErrorResponse(w, err)

	redirect, decodeErr := h.serializer().Decode(requestBody)
	handleErrorResponse(w, decodeErr)

	saveError := h.redirectService.Save(redirect)
	handleErrorResponse(w, saveError)

	response, encodeError := h.serializer().Encode(redirect)
	handleErrorResponse(w, encodeError)

	handleSuccessResponse(w, response, r.Header.Get("Content-Type"), http.StatusCreated)
}

func (h *handler) serializer() ports.RedirectSerializer {
	return &json2.Redirect{}
}

func handleSuccessResponse(w http.ResponseWriter, data []byte, contentType string, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	handleErrorResponse(w, err)
}

func handleErrorResponse(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

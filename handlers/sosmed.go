package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "wayslink/dto/result"
	sosmeddto "wayslink/dto/sosmed"
	"wayslink/models"
	"wayslink/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerSosmed struct {
	SosmedRepository repositories.SosmedRepository
}

func HandlerSosmed(SosmedRepository repositories.SosmedRepository) *handlerSosmed {
	return &handlerSosmed{SosmedRepository}
}

func (h *handlerSosmed) CreateSosmed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get file name
	dataContext := r.Context().Value("dataFile")
	filename := dataContext.(string)

	linkID, _ := strconv.Atoi(r.FormValue("link_id"))

	request := sosmeddto.SosmedRequest{
		LinkID:      linkID,
		TitleSosmed: r.FormValue("title_sosmed"),
		Url:         r.FormValue("url"),
		Image:       r.FormValue("file"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	filePath := os.Getenv("FILE_PATH")

	sosmed := models.Sosmed{
		LinkID:      request.LinkID,
		TitleSosmed: request.TitleSosmed,
		Url:         request.Url,
		Image:       filename,
	}

	data, err := h.SosmedRepository.CreateSosmed(sosmed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	sosmedResponse := sosmeddto.SosmedResponse{
		LinkID:      data.LinkID,
		TitleSosmed: data.TitleSosmed,
		Url:         data.Url,
		Image:       filePath + data.Image,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: sosmedResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerSosmed) FindSosmedsByLinkID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	link_id, _ := strconv.Atoi(mux.Vars(r)["link_id"])

	sosmeds, err := h.SosmedRepository.FindSosmedsByLinkID(link_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	filePath := os.Getenv("FILE_PATH")
	sosmedResponse := make([]sosmeddto.SosmedResponse, 0)

	for _, sosmed := range sosmeds {
		sosmedResponse = append(sosmedResponse, sosmeddto.SosmedResponse{
			ID:          sosmed.ID,
			LinkID:      sosmed.LinkID,
			TitleSosmed: sosmed.TitleSosmed,
			Url:         sosmed.Url,
			Image:       filePath + sosmed.Image,
		})
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: sosmedResponse}
	json.NewEncoder(w).Encode(response)
}

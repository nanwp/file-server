package rest

import (
	"file-server/lib/config"
	"file-server/lib/helper"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	cfg config.Config
}

func NewHandler(
	cfg config.Config,
) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// get file from request
	// save file to disk
	// return file path

	file, header, err := r.FormFile("file")
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		helper.WriteResponse(w, err, nil)
		return
	}
	defer file.Close()

	tipe := r.FormValue("tipe")
	if tipe == "" {
		helper.WriteResponse(w, helper.NewErrBadRequest("Tipe is required"), nil)
		return
	}

	if tipe != "IMAGE" && tipe != "DOCUMENT" {
		helper.WriteResponse(w, helper.NewErrBadRequest("Tipe must be IMAGE or DOCUMENT"), nil)
		return
	}
	// validate content type (image/jped, image/png, application/pdf)
	if header.Header.Get("Content-Type") != "image/jpeg" && header.Header.Get("Content-Type") != "image/png" && header.Header.Get("Content-Type") != "application/pdf" {
		helper.WriteResponse(w, helper.NewErrBadRequest("Content type not allowed"), nil)
		return
	}

	fileName := time.Now().Format("20060102150405") + "-" + header.Filename

	f, err := os.Create("./storage/" + tipe + "/" + fileName)
	if err != nil {
		err = os.MkdirAll("./storage/"+tipe, os.ModePerm)
		if err != nil {
			helper.WriteResponse(w, err, nil)
			return
		}

		f, err = os.Create("./storage/" + tipe + "/" + header.Filename)
		if err != nil {
			helper.WriteResponse(w, err, nil)
			return
		}
	}

	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.WriteResponse(w, nil, map[string]string{"url": h.cfg.APPConfig.BaseURL + "/public/file/storage/" + tipe + "/" + fileName})
}

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	tipe := vars["tipe"]

	log.Println("Get file", filename)

	http.ServeFile(w, r, "./storage/"+tipe+"/"+filename)
}

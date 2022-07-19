package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/abdalrazzak/gin-golang-test/api/auth"
	"github.com/abdalrazzak/gin-golang-test/api/models"
	"github.com/abdalrazzak/gin-golang-test/api/responses"
	"github.com/abdalrazzak/gin-golang-test/api/utils/formaterror"
)

func (server *Server) CreateFile(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	file := models.File{}
	err = json.Unmarshal(body, &file)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	file.Prepare()
	err = file.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != file.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	fileCreated, err := file.SaveFile(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, fileCreated.ID))
	responses.JSON(w, http.StatusCreated, fileCreated)
}

func (server *Server) GetFiles(w http.ResponseWriter, r *http.Request) {

	file := models.File{}

	files, err := file.FindAllFiles(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, files)
}
  

func (server *Server) DeleteFile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	file := models.File{}

	fid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = file.DeleteFile(server.DB, fid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", fid))
	responses.JSON(w, http.StatusNoContent, "")
}

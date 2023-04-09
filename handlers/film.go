package handlers

import (
	filmdto "dumbmerch/dto/film"
	dto "dumbmerch/dto/result"
	"fmt"

	"dumbmerch/models"

	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilm(c echo.Context) error {
	films, err := h.FilmRepository.FindFilm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	for i, p := range films {
		films[i].ThumbnailFilm = path_file + p.ThumbnailFilm
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerFilm) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var film models.Film
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	film.ThumbnailFilm = path_file + film.ThumbnailFilm
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerFilm) CreateFilm(c echo.Context) error {
	// request := new(filmdto.CreateFilmRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	year, _ := strconv.Atoi(c.FormValue("year"))
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))

	request := filmdto.CreateFilmRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		Year:          year,
		CategoryID:    category_id,
		Description:   c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// data form pattern submit to pattern entity db film
	films := models.Film{
		Title:         request.Title,
		ThumbnailFilm: request.ThumbnailFilm,
		Year:          request.Year,
		Category:      request.Category,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
	}

	data, err := h.FilmRepository.CreateFilm(films)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertFilmResponse(data)})
}

func (h *handlerFilm) UpdateFilm(c echo.Context) error {
	request := new(filmdto.UpdateFilmRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	film, err := h.FilmRepository.GetFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}
	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}
	if request.Year != 0 {
		film.Year = request.Year
	}
	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}
	if request.Description != "" {
		film.Description = request.Description
	}

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerFilm) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FilmRepository.DeleteFilm(film, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertDeleteFilm(data)})
}

func convertFilmResponse(u models.Film) filmdto.CreateFilmResponse {
	return filmdto.CreateFilmResponse{
		ID:            u.ID,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		Year:          u.Year,
		Category:      u.Category,
		CategoryID:    u.CategoryID,
		Description:   u.Description,
	}
}

func convertDeleteFilm(u models.Film) filmdto.FilmDeleteResponse {
	return filmdto.FilmDeleteResponse{
		ID: u.ID,
	}
}

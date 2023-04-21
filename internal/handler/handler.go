package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/controller"
	"github.com/jvhab/Redis-crud/internal/model"
	"net/http"
)

type Handler struct {
	ctr    *controller.Controller
	router *gin.Engine
}

func NewHandler(ctr *controller.Controller, router *gin.Engine) *Handler {
	return &Handler{
		ctr:    ctr,
		router: router,
	}
}

func (h *Handler) CreateNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		news := &model.News{}
		err := json.NewDecoder(c.Request.Body).Decode(news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		id, err := h.ctr.Create(c.Request.Context(), news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		c.JSON(http.StatusCreated, id.String())
	}
}

func (h *Handler) UpdateNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		news := &model.News{}
		err = json.NewDecoder(c.Request.Body).Decode(news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update1")
			return
		}
		id := c.Params.ByName("id")
		news.Id, err = uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update2")
			return
		}
		err = h.ctr.Update(c.Request.Context(), news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update3")
			return
		}
		c.JSON(http.StatusOK, "Update "+id)
	}
}

func (h *Handler) GetNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		idTemp := c.Params.ByName("id")
		id, err := uuid.Parse(idTemp)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont parse id")
			return
		}
		result, err := h.ctr.Get(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont have id")
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func (h *Handler) DeleteNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		idTemp := c.Params.ByName("id")
		id, err := uuid.Parse(idTemp)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont parse id")
			return
		}
		err = h.ctr.Delete(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont delete")
			return
		}
		c.String(http.StatusOK, "delete "+idTemp)
	}
}

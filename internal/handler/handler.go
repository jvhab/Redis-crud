package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

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
		news := &model.News{}
		err := json.NewDecoder(c.Request.Body).Decode(news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update")
			return
		}
		id := c.Params.ByName("id")
		news.Id, err = uuid.FromBytes([]byte(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update")
			return
		}
		err = h.ctr.Update(c.Request.Context(), news)
		if err != nil {
			c.JSON(http.StatusBadRequest, "dont update")
			return
		}
		c.JSON(http.StatusOK, "Update")
	}
}

/*
Create(context.Context, *model.News) (uuid.UUID, error)
Update(context.Context, *model.News) error
Get(context.Context, uuid.UUID) (*model.News, error)
Delete(context.Context, uuid.UUID) error
*/

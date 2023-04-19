package handler

func (h *Handler) RegisterRouter() {
	h.router.POST("/create", h.CreateNews())
	h.router.PATCH("/news/:id", h.UpdateNews())
}

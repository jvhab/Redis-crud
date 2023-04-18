package handler

func (h *Handler) RegisterRouter() {
	h.router.POST("/create", h.CreateNews())
}

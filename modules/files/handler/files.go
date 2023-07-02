package handler

import (
	"golang_drive/core/response"
	"golang_drive/modules/files/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type filesHandler struct {
	service service.FilesService
}

func NewFilesHandler(service service.FilesService) *filesHandler {
	return &filesHandler{service}
}

func (h *filesHandler) GetAllFiles(c *gin.Context) {
	files, err := h.service.GetAllFiles()

	if err != nil {
		response := response.ApiResponse("Failed get files data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := response.ApiResponse("Success get files data", http.StatusOK, "error", files)
	c.JSON(http.StatusOK, response)
}

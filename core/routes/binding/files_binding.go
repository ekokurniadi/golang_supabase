package binding

import (
	"golang_drive/modules/files/handler"
	"golang_drive/modules/files/repository"
	"golang_drive/modules/files/service"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

type filesHandlerBinding struct {
	supabase *supabase.Client
	group    *gin.RouterGroup
}

func NewFilesHandlerBinding(supabase *supabase.Client, group *gin.RouterGroup) *filesHandlerBinding {
	return &filesHandlerBinding{supabase, group}
}

func (b *filesHandlerBinding) Run() {
	repo := repository.NewFilesRepository(b.supabase)
	service := service.NewFilesService(repo)
	handler := handler.NewFilesHandler(service)

	b.group.GET("/files", handler.GetAllFiles)
}

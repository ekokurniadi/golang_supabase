package routes

import (
	"golang_drive/core/routes/binding"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

type routes struct {
	supabaseClient *supabase.Client
}

func NewRoutes(supabaseClient *supabase.Client) *routes {
	return &routes{supabaseClient}
}

func (r *routes) RunAppRouter() {
	router := gin.Default()
	api := router.Group("/api/v1")

	binding.NewFilesHandlerBinding(r.supabaseClient, api).Run()

	router.Run()
}

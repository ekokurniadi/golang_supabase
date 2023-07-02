package main

import (
	"golang_drive/core/database"
	"golang_drive/core/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseInstance := database.NewSupabase()

	routes.NewRoutes(supabaseInstance).RunAppRouter()

}

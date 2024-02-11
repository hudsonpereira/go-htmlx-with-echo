package main

import (
	"database/sql"
	"myapp/internal/database"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Api struct {
	DB *database.Queries
}

func main() {
	e := echo.New()

	godotenv.Load(".env")

	dbUrl := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		e.Logger.Fatal("cant connect")
	}

	queries := database.New(conn)

	api := Api{
		DB: queries,
	}

	if err != nil {
		e.Logger.Fatal("cant create queries")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", api.handlerGetUsers)
	e.GET("/users/:ID", api.handlerGetUser)
	e.POST("/users", api.handlerCreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"encoding/json"
	"fmt"
	"myapp/internal/database"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (api *Api) handlerGetUsers(c echo.Context) error {
	users, err := api.DB.GetUsers(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Printf("%v", users)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	if users != nil {
		return c.JSON(http.StatusOK, users)
	}

	return c.JSON(http.StatusOK, []database.User{})
}

func (api *Api) handlerCreateUser(c echo.Context) error {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}

	err := json.NewDecoder(c.Request().Body).Decode(&params)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := api.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (api *Api) handlerGetUser(c echo.Context) error {
	id := c.Param("ID")

	parsedId, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := api.DB.GetUser(c.Request().Context(), parsedId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

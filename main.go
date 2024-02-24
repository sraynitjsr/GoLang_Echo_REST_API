package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]*User{}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.ID = len(users) + 1
	users[user.ID] = user
	return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	user, ok := users[userID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	user, ok := users[userID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	updatedUser := new(User)
	if err := c.Bind(updatedUser); err != nil {
		return err
	}
	user.Name = updatedUser.Name
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	_, ok := users[userID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	delete(users, userID)
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	userList := []*User{}
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(http.StatusOK, userList)
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Building REST API Using Echo Framework")
}

func main() {
	e := echo.New()

	e.GET("/home", home)
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

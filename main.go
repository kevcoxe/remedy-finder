package main

import (
	"fmt"
	"net/http"
	"remedy-filder/db"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	client, err := db.NewSQLiteClient()
	if err != nil {
		panic(fmt.Sprintf("Could not create db client: %v", err))
	}
	defer client.CloseDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Home Page")
	})

	e.GET("/remedies", func(c echo.Context) error {
		remedies := client.GetRemedies()

		response := ""
		for _, remedy := range remedies {
			response += remedy.Name + "\n"
		}

		return c.String(http.StatusOK, response)
	})

	e.GET("/api/v1/remedies", func(c echo.Context) error {
		r := client.GetRemedies()

		fmt.Printf("Remedies: %v\n", r)

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/api/v1/remedies/:id", func(c echo.Context) error {
		id := c.Param("id")

		r, err := client.GetRemedyById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		return c.JSON(http.StatusOK, r)
	})

	e.PUT("/api/v1/remedies", func(c echo.Context) error {
		// Get name and description
		name := c.FormValue("name")
		description := c.FormValue("description")

		r, err := client.CreateRemedy(name, description)
		if err != nil {
			fmt.Printf("error creating remedy: %v\n", err)
			return c.JSON(http.StatusNotFound, nil)
		}

		fmt.Printf("Remedy: %v\n", r)

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello %v!", name))
	})

	e.Logger.Fatal(e.Start(":1323"))
}

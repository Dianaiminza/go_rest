package controller

import (
	"net/http"
	"hello/config"
	"hello/model"
	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	b := new(model.Book)
	db := config.DB()

	// Binding data
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to bind request data",
			"error":   err.Error(),
		})
	}

	book := &model.Book{
		Id: 1,
		Name:        b.Name,
		Description: b.Description,
	}

	// Creating the book record
	if err := db.Create(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create book",
			"error":   err.Error(),
		})
	}

	response := map[string]interface{}{
		"message": "Book created successfully",
		"data":    book,
	}

	return c.JSON(http.StatusOK, response)
	
}

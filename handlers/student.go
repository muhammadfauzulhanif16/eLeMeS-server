package handlers

import (
	"eLeMeS/inputs"
	"eLeMeS/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type studentHandler struct {
	s services.Student
}

func StudentHandler(s services.Student) *studentHandler {
	return &studentHandler{s}
}

func (h *studentHandler) Add(c *gin.Context) {
	var input inputs.AddStudent

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	data, err := h.s.Add(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, data)
}

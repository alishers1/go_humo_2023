package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var data Calculator

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := data.FirstNum + data.SecondNum
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Subtract(c *gin.Context) {
	var data Calculator

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := data.FirstNum - data.SecondNum

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Multiply(c *gin.Context) {
	var data Calculator

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := data.FirstNum * data.SecondNum

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Divide(c *gin.Context) {
	var data Calculator

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := data.FirstNum / data.SecondNum

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

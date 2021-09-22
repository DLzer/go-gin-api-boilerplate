package controllers

import (
	"github.com/gin-gonic/gin"
)

// Response object as HTTP response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTPRes normalize HTTP Response format
func HTTPRes(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Status:  httpCode,
		Message: msg,
		Data:    data,
	})
}

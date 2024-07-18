package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (o *OrderController) GetById(c *gin.Context) {
	fmt.Println("Create")
}

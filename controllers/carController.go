package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var carDatas []Car

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		err = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf(`c%d`, len(carDatas)+1)

	carDatas = append(carDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})

}

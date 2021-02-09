package controller

import {
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"net/http"
	"time"
}

func (idb *InDB) CreateCustomer(ctx *gin.Context){
	var (
		costumer model.Contract_Accounts
	)

	err := ctx.ShouldBindJSON(&costumer)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}
}


package controller

import (
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"net/http"
	"time"
	"github.com/chandrartw/cobacrud/util"
	"github.com/chandrartw/cobacrud/model"
)

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
	util.ResponseSuccess(ctx, http.StatusOK, costumer)
}

func (idb *InDB) GetCostumer(ctx *gin.Context) {
	var (
		costumer   model.Contract_Accounts
	)
	id := ctx.Param("CA_No")
	err := idb.DB.Where("CA_No = ?", id).First(&customer).Error

	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	util.ResponseSuccess(ctx, http.StatusOK, user)
}

func (idb *InDB) GetAllCustomer(ctx *gin.Context) {

	customers := []model.Contract_Accounts{}
	_ = idb.DB.Find(&customers).Error
	for i, _ := range customers {
		idb.DB.Model(customers[i])
	}
	if len(customers) <= 0 {
		util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "No Record")
	} else {
		util.ResponseSuccess(ctx, http.StatusOK, users)
	}

}

func (idb *InDB) DeleteUser(ctx *gin.Context) {
	var (
	      customer model.Contract_Accounts
	)
	id := ctx.Param("CA_No")

	err := idb.DB.Where("CA_No = ?", id).First(&customer).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	err = idb.DB.Unscoped().Delete(&customer).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Find Id User")
		ctx.Abort()
		return
	}
	if err != nil {
		util.ResponseError(ctx, http.StatusBadRequest, err.Error(), "Error Delete User")
		ctx.Abort()
		return
	}

	util.ResponseSuccessCustomMessage(ctx, http.StatusOK, "Success Deleted")
}

func (idb *InDB) UpdateUser(ctx *gin.Context) {
	ca_no := ctx.Query("CA_No")

	var (
		customer   model.Contract_Accounts
		newcostumer model.Contract_Accounts
	)

	err := idb.DB.First(&customer, ca_no).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "ID Not Found")
		ctx.Abort()
		return
	}

	err = ctx.ShouldBindJSON(&newcostumer)
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Please Check Your Data")
		ctx.Abort()
		return
	}

	err = idb.DB.Model(&customer).Updates(newcustomer).Error
	if err != nil {
		util.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error(), "Update Failed")
		ctx.Abort()
		return
	}

	util.ResponseSuccess(ctx, http.StatusOK, user)

}



package model

import (
	"time"
	 "github.com/jinzhu/gorm"
)

type Contract_Accounts struct {
	gorm.model
	CA_No string `json:"ca" binding:"required" gorm:"type:varchar(20)"`
	BP_No string `json:"bp" binding:"required" gorm:"type:varchar(20)"`
	Created_Date string `json:"created_date" binding:"required" gorm:"default:CURRENT_TIMESTAMP"`
	Created_By	string `json:"createdby" binding:"required" gorm:"type:varchar(20)"`
	Valid_To	time.Time `json:"validto" binding:"required"`
	CA_Type string `json:"ca_type" binding:"required" gorm:"type:varchar(20)"`
	CA_Name string `json:"ca_name" binding:"required" gorm:"type:varchar(20)"`
	CCA string `json:"cca" binding:"required" gorm:"type:varchar(20)"`
	Business_Area string `json:"business_area" binding:"required" gorm:"type:varchar(20)"`
	BP_Relation string `json:"bp_relation" binding:"required" gorm:"type:varchar(20)"`
	Trading_Partner string `json:"trading_partner" binding:"required" gorm:"type:varchar(20)"`
	Currency string `json:"currency" binding:"required" gorm:"type:varchar(20)"`
	Auth_Group string `json:"auth_group" binding:"required" gorm:"type:varchar(20)"`
	Refno string `json:"refno" binding:"required" gorm:"type:varchar(20)"`
	Pay_Terms string `json:"pay_terms"  gorm:"type:varchar(20)"`
	Tolerance_Grp string `json:"tolerance_grp"  gorm:"type:varchar(20)"`
	Clearing_Cat string `json:"clearing_cat"  gorm:"type:varchar(20)"`
	Acc_Determ string `json:"acc_determ" gorm:"type:varchar(20)"`
	Plan_Grp string `json:"plan_grp"  gorm:"type:varchar(20)"`
	Interest_Key string `json:"interest_key"  gorm:"type:varchar(20)"`
	Kode_Sentral string `json:"kode_sentral"  gorm:"type:varchar(20)"`
	Kode_Catel string `json:"kode_catel"  gorm:"type:varchar(20)"`
	Payer_BP string `json:"payer_bp"  gorm:"type:varchar(20)"`
	Payer_CA string `json:"payer_ca"  gorm:"type:varchar(20)"`
	Isolir_Free bool `json:"isolir_free" gorm:"default:false"`
	Vat_Free bool `json:"vat_free" gorm:"default:false""`
	Stamp_Duty bool `json:"stamp_duty"  gorm:"default:false"`
	PPN_5PCT bool `json:"ppn_5pct"  gorm:"default:false"`
	Dunning_Grp bool `json:"dunning_grp"  gorm:"default:false"`

}
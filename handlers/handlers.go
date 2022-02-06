package handlers

import "github.com/gin-gonic/gin"

type IHandler interface {
	BindData(*bool)
	GetAllRecordsHandler(ctx *gin.Context)
}

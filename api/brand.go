package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/category/models"
	"github.com/provider-go/pkg/ioput"
)

func CreateCategory(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)

	parentId := ioput.ParamToInt32(json["parentId"])
	categoryName := ioput.ParamToString(json["categoryName"])
	icon := ioput.ParamToString(json["icon"])
	pic := ioput.ParamToString(json["pic"])
	brief := ioput.ParamToString(json["brief"])
	err := models.Createcategory(parentId, categoryName, icon, pic, brief)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func DeleteCategory(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	err := models.Deletecategory(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, "success")
	}
}

func ListCategory(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	pageSize := ioput.ParamToInt(json["pageSize"])
	pageNum := ioput.ParamToInt(json["pageNum"])
	list, total, err := models.Listcategory(pageSize, pageNum)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		ioput.ReturnSuccessResponse(ctx, res)
	}
}

func ViewCategory(ctx *gin.Context) {
	json := make(map[string]interface{})
	_ = ctx.BindJSON(&json)
	id := ioput.ParamToInt32(json["id"])
	row, err := models.Viewcategory(id)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, row)
	}
}

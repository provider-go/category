package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/category/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// category 表操作
		Router.POST("add", api.CreateCategory)
		Router.POST("delete", api.DeleteCategory)
		Router.POST("list", api.ListCategory)
		Router.POST("view", api.ViewCategory)
	}
}

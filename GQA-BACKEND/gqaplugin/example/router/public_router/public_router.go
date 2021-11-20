package public_router

import (
	"gin-quasar-admin/gqaplugin/example/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	{
		publicGroup.GET("news-list", public_api.GetNews)
	}
}
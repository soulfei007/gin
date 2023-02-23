package routers

import (
	"gin-learn/pkg/setting"
	"gin-learn/routers/api"
	v1 "gin-learn/routers/api/v1"
	"net/http"

	"gin-learn/pkg/upload"

	"github.com/gin-gonic/gin"
)

//路由设置
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //用中间件、
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/auth", api.GetAuth) //客户端验证获取token
	r.POST("/upload", api.UploadImage)
	apiv1 := r.Group("/api/v1") //分组路由
	//apiv1.Use(jwt.JWT())        //分组路由中间件验证token
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除文章
		apiv1.DELETE("/articles/:id")
	}

	return r
}

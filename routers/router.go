package routers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"

    "github.com/weekndCN/shorturl/routers/api"
    "github.com/weekndCN/shorturl/pkg/setting"
    "github.com/weekndCN/shorturl/models"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    gin.SetMode(setting.ServerSetting.RunMode)

    r.GET("/:shorten_id", func(c *gin.Context){
        shorten := c.Param("shorten_id")
        fmt.Println(shorten)
        _, lengthen := models.GetLengthen(shorten)
        if lengthen != "" {
            c.Redirect(http.StatusMovedPermanently,lengthen)
        }
        c.Redirect(http.StatusMovedPermanently,"https://www.jianshu.com/")
    })

	r.POST("/j/shorten", api.GetShorten)
	r.POST("/j/lengthen", api.GetLengthen)
    return r
}

package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "github.com/weekndCN/shorturl/models"
    "github.com/weekndCN/shorturl/pkg/e"
)



func GetLengthen(c *gin.Context) {
    shorten := c.PostForm("shorten")
    data := make(map[string]interface{})
    code := e.SUCCESS
    // 获取数据库的主键ID
    id, lengthen := models.GetLengthen(shorten)
    data["id"] = id
    data["lengthen"] = lengthen
    data["shorten"] = shorten
    // 返回请求结果
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}


func GetShorten(c *gin.Context) {
    lengthen := c.PostForm("lengthen")
    fmt.Println(lengthen)
    data := make(map[string]interface{})
    code := e.SUCCESS
    // 获取数据库的主键ID
    id, shorten := models.GetShorten(lengthen)
	// id=0 表示不存在，新建一个
	if id == 0 {
		id, lengthen, shorten = models.AddUrl(lengthen)
	}
    data["id"] = id
    data["lengthen"] = lengthen
    data["shorten"] = shorten
    // 返回请求结果
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

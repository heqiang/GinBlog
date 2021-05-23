package v1

import (
	"GinDemo/src/models"
	"GinDemo/src/pkg/e"
	"GinDemo/src/pkg/setting"
	"GinDemo/src/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章的标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}

	code := e.SUCCESS
	pageName := util.GetPage(c)
	pageSize := setting.PageSize
	fmt.Println("**********")
	fmt.Println(maps)
	fmt.Println(pageName)
	fmt.Println(pageSize)
	data["lists"] = models.GetTags(pageName, pageSize, maps)
	//
	//data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//增加文章标签
func  AddTags(c *gin.Context)  {
	c.String(200,"增加文章的标签")
}
//修改文章标签
func  EditTags(c *gin.Context)  {
	c.String(200,"修改文章的标签")
}
//删除文章标签
func  DeleteTags(c *gin.Context)  {
	c.String(200,"删除文章的标签")
}
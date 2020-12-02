package v1

import (
	"go-gin-app/models2"
	"go-gin-app/pkg/e"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	// data["lists"] = models.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetUserById 根据ID获取用户信息
func GetUserById(c *gin.Context) {
	code := e.SUCCESS

	id := c.Param("id")
	data := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id不能为空")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	data["user"] = models2.GetUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

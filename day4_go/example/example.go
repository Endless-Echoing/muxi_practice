package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetJSON 获取JSON
func GetJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

// GetSomeJSON 获取JSON
func GetSomeJSON(c *gin.Context) {
	var user User
	user.Name = "John"
	user.Age = 20
	c.JSON(http.StatusOK, user)
}

// PostJSON1 结构体方式传递JSON数据
func PostJSON1(c *gin.Context) {
	var user User
	// 将请求体中的JSON数据绑定到user变量中，可以不传，不传则为零值
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// PostJSON2 表单方式传递数据，接收表单字段
func PostJSON2(c *gin.Context) {
	// DefaultPostForm 获取POST表单参数，如果没有则使用默认值
	username := c.DefaultPostForm("username", "anonymous")
	//username := c.PostForm("username")
	address := c.PostForm("address")
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

// PostJSON3 查询参数方式传递数据，接收URL参数
func PostJSON3(c *gin.Context) {
	// DefaultQuery 获取URL中的查询参数，如果没有则使用默认值
	// 也可以用 c.Query("username")
	username := c.DefaultQuery("username", "默认用户")
	//username := c.Query("username")
	address := c.Query("address")

	message := fmt.Sprintf("Hello %s, your address is %s", username, address)
	// 返回json
	c.JSON(http.StatusOK, gin.H{
		"message":  message,
		"username": username,
		"address":  address,
	})
}

// PostJSON4 路径参数方式传递数据
func PostJSON4(c *gin.Context) {
	username := c.Param("username")
	address := c.Param("address")
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var user = User{Name: "John", Age: 20}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ShouldBind 结构体/表单/查询参数方式绑定数据
func ShouldBind(c *gin.Context) {
	var login Login
	// ShouldBind()会根据请求的Content-Type自动选择绑定器
	// 如果是 GET 请求，只使用 Form 绑定引擎(query)
	// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form(form-data)
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, login)
}

func main() {
	// 创建一个默认的路由
	r := gin.Default()
	// 注册路由和处理函数
	r.GET("/JSON", GetJSON)
	r.GET("/someJSON", GetSomeJSON)
	r.POST("/JSON1", PostJSON1)
	r.POST("/JSON2", PostJSON2)
	r.POST("/JSON3", PostJSON3)
	r.POST("/JSON4/:username/:address", PostJSON4)
	r.PUT("/user", UpdateUser)
	r.POST("/ShouldBind", ShouldBind)
	r.GET("/ShouldBind", ShouldBind)
	// 默认监听本地地址的8080端口和错误处理
	r.Run(":8080")
}

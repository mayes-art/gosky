package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type apiResult struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	DataGrid interface{} `json:"dataGrid"`
}

var users = []User{
	{ID: 1, Name: "Jun", Email: "jun@test.com"},
	{ID: 2, Name: "Yue", Email: "yue@test.com"},
	{ID: 3, Name: "Kay", Email: "kay@test.com"},
}

func successResponse(data interface{}) apiResult {
	return apiResult{
		Status:   1,
		Message:  "OK",
		DataGrid: data,
	}

	// c.JSON(http.StatusOK, r)
}

func failedResponse(c *gin.Context, msg string) {
	r := apiResult{
		Status:   1,
		Message:  msg,
		DataGrid: nil,
	}

	c.JSON(http.StatusOK, r)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.IndentedJSON(http.StatusBadRequest, gin.H{"status": -1, "message": "Parameter is invailid!"})
		failedResponse(c, "Parameter is invailid!")
	}

	for _, u := range users {
		if u.ID == id {
			// c.IndentedJSON(http.StatusOK, u)

			c.JSON(http.StatusOK, successResponse(u))

			// successResponse(c, u)
			return
		}
	}

	// c.IndentedJSON(http.StatusNotFound, gin.H{"status": 0, "message": "No data"})
	failedResponse(c, "No data")
}

func addUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func main() {
	// fmt.Println("Helloooooooo~")

	l := logrus.New()
	l.Info("log init")

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "message": "Server is running!"})
	})

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", getUsers)
			users.GET("/:id", getUser)
			users.POST("/", addUser)
		}

	}

	r.Run(":999")
}

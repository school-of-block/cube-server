package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

  api := r.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }

  api.GET("/users", GetAllUsersHandler)
  api.GET("/users/:userID", GetUserHandler)


	r.Run(":8080")
}

func GetAllUsersHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"Here are all the users...",
  })
}

func GetUserHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"Here is the user data",
  })
}


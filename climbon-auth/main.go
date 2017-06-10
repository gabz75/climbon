package main

import (
  "github.com/gin-gonic/gin"
  "gopkg.in/appleboy/gin-jwt.v2"
  "time"
)

func main() {
  r := gin.Default()

  authMiddleware := &jwt.GinJWTMiddleware{
    Realm:      "test zone",
    Key:        []byte("secret key"),
    Timeout:    time.Hour,
    MaxRefresh: time.Hour * 24,
    Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
      if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
        return userId, true
      }

      return userId, false
    },
    Authorizator: func(userId string, c *gin.Context) bool {
      if userId == "admin" {
        return true
      }

      return false
    },
    Unauthorized: func(c *gin.Context, code int, message string) {
      c.JSON(code, gin.H{
        "code":    code,
        "message": message,
      })
    },
    TokenLookup: "header:Authorization",
  }

  r.POST("/login", authMiddleware.LoginHandler)

  authMiddleware.MiddlewareFunc()

  token := r.Group("/token")
  token.Use(authMiddleware.MiddlewareFunc())
  {
    token.GET("/valid", func(c *gin.Context) {
      c.JSON(204, gin.H{})
    })
    token.GET("/refresh", authMiddleware.RefreshHandler)
  }

  r.Run(":8080")
}


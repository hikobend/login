package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.POST("/login", func(c *gin.Context) {
		// セッションの作成
		session := sessions.Default(c)
		session.Set("loginUser", c.PostForm("userId"))
		session.Save()
		c.String(http.StatusOK, "ログイン完了")
	})
	r.GET("/logout", func(c *gin.Context) {
		// セッションの破棄
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.String(http.StatusOK, "ログアウトしました")
	})
	r.Run() // listen and server on 0.0.0.0:8080
}

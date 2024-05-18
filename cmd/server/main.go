package main

import "github.com/gin-gonic/gin"

type Blog struct {
	Title string `json:"title"`
}

func main() {
	r := gin.New()
	r.Static("/static", "../../static")

	r.GET("/", func(c *gin.Context) {
		c.File("../../static/index.html")
	})

	r.GET("/api/blogs", func(c *gin.Context) {
		c.JSON(200, []Blog{
			{Title: "Blog 1"},
			{Title: "Blog 2"},
		})
	})

	r.Run(":8080")
}

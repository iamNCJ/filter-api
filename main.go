package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	err := filter.LoadWordDict("dict.txt")
	if err != nil {
		fmt.Println("Error loading filter list! Try loading a network version")
		_ = filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		body := string(buf[0:n])
		fmt.Print(body)
		res, illegalWord := filter.Validate(body)
		filtered := body
		if !res {
			filtered = filter.Filter(body)
		}
		c.JSON(200, gin.H{
			"valid":       res,
			"illegalWord": illegalWord,
			"filtered":    filtered,
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

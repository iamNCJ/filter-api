package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
)

func main() {
	// filter
	filter := sensitive.New()
	err := filter.LoadWordDict("dict.txt")
	if err != nil {
		fmt.Println("Error loading filter list! Try loading a network version")
		_ = filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	}

	// gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		body := c.Query("string")
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

	// run
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

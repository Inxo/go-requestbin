package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		app, okApp := c.GetQuery("app")
		api, okApi := c.GetQuery("api")
		if okApp && okApi {
			if Apps[app] == api {
				data, err := c.GetRawData()
				if err == nil {
					writeInFile(app, data)
				} else {
					println("error")
				}
			}
		}
	})

	r.Run(":8004")
}

func writeInFile(app string, data []byte) {
	filename := "logs/" + app + ".log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
}

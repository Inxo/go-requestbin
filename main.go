package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	 println("Work directory:" + dir)
	loadEnv()
	r := gin.Default()
	r.POST("/requestbin/", func(c *gin.Context) {
		app, okApp := c.GetQuery("app")
		api, okApi := c.GetQuery("api")
		if okApp && okApi {
			if os.Getenv(app) == api {
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
	directory := os.Getenv("wd")
	filename := directory + "/logs/" + app + ".log"
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
}

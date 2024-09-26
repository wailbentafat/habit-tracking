package main

import (
	"log"

	"habit/db"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"habit/controlers"

)


func main() {
    db.SetupDatabase()
    r := gin.Default()
	c := cron.New()
    c.AddFunc("0 0 * * *", func() {
		log.Println("Checking for missed habits...")
		controlers.LogMissedProgress()
	})
    c.Start()
	defer c.Stop()
   
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

package controlers

import (
	"habit/db"
	"habit/models"
	"log"
	"time"

)
func LogMissedProgress(){

	var habits[]models.Habit
	if err:=db.DB.Find(&habits).Error;err!=nil{
		log.Println("Failed to get habits:", err)
		return
	}
	for _,habits :=range habits{
		var progress models.Progres
		today:=time.Now().Truncate(24 * time.Hour)
		if err:=db.DB.Where("habit_id=? AND date=?",habits.ID,today).First(&progress).Error;err!=nil{
			progress:=models.Progres{
				Habit_id:habits.ID,
				Date:today,
				Status:"missed",
			}
			if err:=db.DB.Create(&progress).Error;err!=nil{
				log.Println("Failed to log missed progress:", err)
				return
			}
		}
	}
log.Println("Missed progress logged successfully")
}
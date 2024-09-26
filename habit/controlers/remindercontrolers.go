package controlers

import (
	"habit/db"
	"habit/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func CraeteReminder(c*gin.Context){
	type reminderinput struct {
	Habit_id uint `json:"habitid"`
	Date time.Time `json:"date"`}

	var input reminderinput
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return}
	reminder:=models.Reminder{
		Habit_id:input.Habit_id,
		Date:input.Date,
		Status:"pending",
	}	

	if err:=db.DB.Create(&reminder).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}	


	c.JSON(http.StatusOK, gin.H{"message": "Reminder created"})
}

func GetReminders(c*gin.Context){
	id:=c.Param("id")
	var reminders[]models.Reminder
	if err:=db.DB.Where("habit_id=?",id).Find(&reminders).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reminders": reminders})
}
func UpdateReminder(c*gin.Context){
	id:=c.Param("id")
	var reminder models.Reminder
	if err:=db.DB.Where("id=?",id).First(&reminder).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	type reminderinput struct {
		Status string `json:"status"`
		Date time.Time `json:"date"`
	} 
	var input reminderinput
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	reminder.Status=input.Status
	reminder.Date=input.Date
	if err:=db.DB.Save(&reminder).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reminder updated"})
}
func DeleteReminder(c*gin.Context){
	id:=c.Param("id")
	var reminder models.Reminder
	if err:=db.DB.Where("id=?",id).First(&reminder).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	if err:=db.DB.Delete(&reminder).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted"})
}
func Getallreminders(c*gin.Context){
	var reminders[]models.Reminder
	if err:=db.DB.Find(&reminders).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reminders": reminders})
}
func SendRemindernotification(c*gin.Context){
 c.Header("content-type","text/event-stream")
 c.Header("cache-control","no-cache")
 c.Header("connection","keep-alive")
 c.Header("access-control-allow-origin","*")
 ticker:=time.NewTicker(1*time.Minute)
 defer ticker.Stop()
 for{
	select{
	case<-ticker.C:
		var reminders[]models.Reminder
		if err:=db.DB.Find(&reminders).Error;err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
			return
		}
		var name string
		for _,reminder:=range reminders{
			if reminder.Date.Before(time.Now()) && reminder.Status=="pending"{
				if err:=db.DB.Model(&models.Habit{}).Select("name").Where("id=?",reminder.Habit_id).First(&name).Error;err!=nil{
					c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
					return
				}	
				}
				c.SSEvent("reminder",gin.H{"name":name})
				
				
			}
		case <-c.Request.Context().Done():
		return
	}	
		}
		
}
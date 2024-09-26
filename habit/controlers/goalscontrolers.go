package controlers

import (
	"habit/db"
	"habit/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func Addgoal(c*gin.Context){
	habit_id:=c.Param("id")
	type Input struct {
		Target int `json:"target"`

	}
	var habit models.Habit

	if err:=db.DB.Where("id=?",habit_id).First(&habit).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	var in Input
	if err:=c.ShouldBindJSON(&in);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	goal:=models.Goals{
		Habit_id:habit.ID,
		Target:in.Target,
		Current:0,
		Createdat:time.Now(),
		UpdatedAt:time.Now(),
	}
	if err:=db.DB.Create(&goal).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Goal created"})
}

func Suivegoal(c*gin.Context){
	id := c.Param("id")
	tx:= db.DB.Begin()
	if err:=tx.Where("id=?",id).First(&models.Goals{}).Error;err!=nil{
		tx.Rollback()
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	var goal models.Goals
	if err:=tx.Model(&models.Goals{}).Where("id=?",id).Select("current,streak,updatedat,target").Scan(&goal).Error;err!=nil{
		tx.Rollback()
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	if goal.Current==goal.Target{
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{"message": "Goal completed"})
		return
	}
	yesterday:=time.Now().AddDate(0,0,-1)

	if goal.UpdatedAt.Before(yesterday){
		goal.Streak=1
	}else{
		goal.Streak++
	}
		

	
	if err := tx.Model(&models.Goals{}).Where("id=?", id).Updates(map[string]interface{}{"current": goal.Current + 1, "updatedat": time.Now(), "streak": goal.Streak}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database"})
		return
	}
	progress := models.Progres{
        Habit_id: goal.Habit_id,
        Date:     time.Now(),
        Status:   "completed", 
    }

	if err := tx.Create(&progress).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"err": "database"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Goal updated"})
}

func Getgoals(c*gin.Context){
	id:=c.Param("id")
	var goals models.Goals
	if err:=db.DB.Where("habit_id=?",id).Select("target,current").Scan(&goals).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	target:=goals.Target
	current:=goals.Current
	percent:=float64(current)/float64(target)*100
	

	c.JSON(http.StatusOK, gin.H{"target":target,"current":current,"percent":percent,"streak":goals.Streak})
}
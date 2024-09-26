package controlers

import (
	"habit/db"
	"habit/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func Addhabit(c*gin.Context){
	usid,ok:=c.Get("user_id")
	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	type habitinput struct {
    Name string `json:"name"`
	Categorie string`json:"categorie"`
	}
	var input habitinput
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var category models.Categorie
	if err:=db.DB.Where("name=?",input.Categorie).Scan(&category).Error;err!=nil{
		c.JSON(http.StatusNoContent,gin.H{"err":"mkcho"})
		return
	}
	id_user, ok := usid.(int)
     if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
}
	habit:=models.Habit{
		Userid:uint(id_user),
		CategoryID:category.ID ,
		Name:input.Name,
		Createdat: time.Now(),
		UpdatedAt: time.Now(),

	}
	if err:=db.DB.Create(&habit).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habit created", "habit": habit})
}
func Gethabits(c*gin.Context){
	usid,ok:=c.Get("user_id")
	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id_user, ok := usid.(int)
	 if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var habits[]models.Habit
	if err:=db.DB.Where("user_id=?",id_user).Find(&habits).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habit created", "habits": habits})
}
func Get_habitbycategories(c*gin.Context){
	categoryname := c.Query("categoryname")
	var category models.Categorie
	if err:=db.DB.Where("name=?",categoryname).First(&category).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	var habits[]models.Habit
	if err:=db.DB.Where("category_id=?",category.ID).Find(&habits).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"habits": habits})
}
func Get_habitbyid(c*gin.Context){
	id:=c.Param("id")
	var habit models.Habit
	if err:=db.DB.Where("id=?",id).First(&habit).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"habit": habit})
}
func Delete_habit(c*gin.Context){
	id:=c.Param("id")
	var habit models.Habit
	if err:=db.DB.Where("id=?",id).First(&habit).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	if err:=db.DB.Delete(&habit).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"err":"database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted"})
}

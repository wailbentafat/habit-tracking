package controlers

import (
	"habit/core/jwt"
	"habit/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"habit/db"
)
func Register(c*gin.Context){
	type RegisterInput struct{
		Username string  `json:"username"`
		Password string `json:"password"`
	}
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := models.User{
		Username: input.Username,
	}
	hash,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "hach creation failed"})
	}
	user.Password=string(hash)
	if err:=db.DB.Create(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": "User creation failed"})
		return
	}
    token,err:=jwt.Createtoken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"err":"glta f token"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"token": token})



	}

	func Login(c *gin.Context) {
		type LoginInput struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
	
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
	
		var user models.User
		if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
	
		// Compare the provided password with the hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
	
		// Create a JWT token
		token, err := jwt.Createtoken(int(user.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
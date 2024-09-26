package middlewear

import (
	"errors"
	"fmt"
	"habit/core/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"habit/models"
)
func Authmiddlwear(db*gorm.DB)gin.HandlerFunc{
	return func(c*gin.Context){
		header:=c.GetHeader("authorization")
		if header==""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		parts:=strings.Split(header," ")
		if len(parts)!=2{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
		}
		token:=parts[1]
		user_id,err:=jwt.Parsing(token)
		if err != nil {
			fmt.Println(err)
			c.Abort()
		}
		var user models.User
		if err:=db.First(&user,user_id).Error;err!=nil{
			if errors.Is(err,gorm.ErrRecordNotFound){
				c.JSON(http.StatusUnauthorized,gin.H{"err":"usernot found"})
				c.Abort()
				return
			}
		}
		c.Set("user_id",user_id)
		c.Next()
        



	
	}
}
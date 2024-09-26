package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


var secretkey=[]byte("secretkeydbdhbcs")
type Claims struct {
    UserId       int       `json:"userId"`
    ExpiryTime   time.Time `json:"exp"`
	jwt.RegisteredClaims  
}

func Createtoken(userid int)(string ,error){
	expiryTime:=time.Now().Add(time.Hour)
	claims := Claims{
        UserId:     userid,
        ExpiryTime: expiryTime,
		RegisteredClaims: jwt.RegisteredClaims{ 
            ExpiresAt: jwt.NewNumericDate(expiryTime),
		}}


	signedToken, err:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims).SignedString(secretkey)
	if err != nil {
		return "",err
	}
	return signedToken,nil

}
func Parsing(token string)(*int,error){
	var claims Claims
	tok,err:=jwt.ParseWithClaims(token,&claims,func(t *jwt.Token) (interface{}, error) {
		if _,ok:=t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,errors.New("signin galat")
		}
		return secretkey,nil
	})
	if err!=nil||!tok.Valid{
		return nil,errors.New("token galat")
	}
	if claims.ExpiresAt.Time.Before(time.Now()){
		return nil,errors.New("token khlass")
	}
	UserId:=claims.UserId
	return &UserId,nil
}

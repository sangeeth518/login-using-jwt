package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sangeeth/jwt-go/initializers"
	"github.com/sangeeth/jwt-go/models"
	"golang.org/x/crypto/bcrypt"
)

//signup
//Get the email and pass

var body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body",
		})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body",
		})
		return
	}

	//Check user
	var user models.User
	initializers.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id or Password"})
	}

	// Validate password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "password dosen't match"})

		return
	}

	// gernerate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

func AdminSignout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Admin Successfully Signed Out",
	})
}

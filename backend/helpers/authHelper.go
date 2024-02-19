package helper

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("UserType")

	err = nil
	if userType != role {
		err = errors.New("unauthorized Access")
		return err
	}
	return err
}
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("userType")
	uid := c.GetString("uid")

	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized Access")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "password is incorrect"
		check = false
	}

	return check, msg
}

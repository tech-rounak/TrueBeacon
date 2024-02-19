package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	database "github.com/tech-rounak/true-beacon/database"
	helper "github.com/tech-rounak/true-beacon/helpers"
	"github.com/tech-rounak/true-beacon/models"
)

var validate = validator.New()
var db = database.InitDB()

func Signup() gin.HandlerFunc {
	return func(gc *gin.Context) {
		fmt.Println("---------------- Signup Handler-------------------------------")
		var user models.User
		// var db = database.GetDB()
		if err := gc.BindJSON(&user); err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)

		if validationErr != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		rows, err := db.Raw("select * from users where user_name = ?", user.UserName).Rows()
		defer rows.Close()

		if err != nil {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var tmpUser models.User
		for rows.Next() {
			err := db.ScanRows(rows, &tmpUser)
			if err != nil {
				gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if tmpUser.UserName != "" {
				gc.JSON(http.StatusInternalServerError, gin.H{"error": "Duplicate User Name . Please choose another username "})
				return
			}

		}
		password := helper.HashPassword(user.Password)
		user.Password = password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		db.Create(&user)

		helper.Success(gc, 201, user, "User SignUp Successfull")
	}
}

func Login() gin.HandlerFunc {
	return func(gc *gin.Context) {
		fmt.Println("-------------------- Login Handler ---------------------------------")

		var user models.User
		var foundUser models.User

		if err := gc.BindJSON(&user); err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rows, err := db.Raw("select * from users where user_name = ?", user.UserName).Rows()
		defer rows.Close()
		if err != nil {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for rows.Next() {
			err := db.ScanRows(rows, &foundUser)
			fmt.Println(foundUser)
			if err != nil {
				gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if foundUser.UserName == "" {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid User Name"})
			return
		}
		passwordIsValid, msg := helper.VerifyPassword(user.Password, foundUser.Password)

		if !passwordIsValid {
			gc.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		token, err := helper.GenerateTokens(foundUser.UserName, foundUser.Name)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		foundUser.Password = ""
		result := map[string]interface{}{
			"user":  foundUser,
			"token": token,
		}

		helper.Success(gc, 201, result, "Loging Successfull")
	}
}

func Profile() gin.HandlerFunc {
	return func(gc *gin.Context) {
		res := map[string]interface{}{
			"user": map[string]interface{}{
				"user_id":   "AB1234",
				"user_type": "individual",
				"email":     "xxxyyy@gmail.com",
				"user_name": "AxAx Bxx",
				"broker":    "ZERODHA",
			},
		}
		helper.Success(gc, 200, res, "Profile data fetched Successfully!!")
	}
}

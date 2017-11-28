package controllers

import (
	"github.com/labstack/echo"
	//"github.com/dgrijalva/jwt-go"
	//"time"
	"net/http"
	//"os/user"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zanjs/go-echo-rest/app/models"
	"github.com/zanjs/go-echo-rest/config"
	"golang.org/x/crypto/bcrypt"
)

var jwtConfig = config.Config.JWT

func PostLogin(c echo.Context) error {

	c.FormParams()

	u := new(models.UserLogin)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Invalid username / password 1",
		})
	}

	fmt.Println(u)

	username := c.FormValue("username")
	password := c.FormValue("password")

	// user, err := models.GetUserByUsername(username)
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	fmt.Println("oassw:", password, username, user.Password)
	fmt.Println("oassw2:", u.Password, user.Password)
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Invalid username / password 2",
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

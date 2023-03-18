package service

import (
	"FoundPlatform/model"
	"FoundPlatform/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetALlUsers returns all users in the database
func GetALlUsers(c *gin.Context) {
	users := repository.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUserByID returns a user by id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user := repository.GetUserByID(id)

	// check if user is found
	if user.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	// check if user input email and password
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email and password are required"})
		return
	}

	// check user input email format with regex
	if !regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`).MatchString(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email format is invalid"})
		return
	}

	// check if user email already exists
	if repository.GetUserByEmail(user.Email).Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	// hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)

	// generate nano id
	user.ID, err = gonanoid.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	// create user
	user = repository.CreateUser(user)
	c.JSON(http.StatusCreated, user)
}

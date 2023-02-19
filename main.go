package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=gotest  port=5432 sslmode=disable")
	if err != nil {

		panic(err)
	}
	defer db.Close()

	// Automigrate the schema
	db.AutoMigrate(&User{})

	r := gin.Default()

	// Create a new user
	r.POST("/user", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		db.Create(&user)
		c.JSON(200, gin.H{"message": "User created successfully"})
	})

	// Get all users
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, gin.H{"data": users})
	})

	// Get a single user by ID
	r.GET("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.First(&user, id)
		c.JSON(200, gin.H{"data": user})
	})

	// Update a user by ID
	r.PUT("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.First(&user, id)
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, gin.H{"message": "User updated successfully"})
	})

	// Delete a user by ID
	r.DELETE("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		db.First(&user, id)
		db.Delete(&user)
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})

	r.Run()
}

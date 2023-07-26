package controllers

import (
	"JWT-Authentication/database"
	"JWT-Authentication/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetEmail(c *gin.Context) (string, string) {
	userInterface, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Cannot get details of user from cookies",
		})
		return "", ""
	}
	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Cannot get details of user from cookies",
		})
		return "", ""
	}
	return user.Name, user.Email
}
func GetNotes(c *gin.Context) {
	name, email := GetEmail(c)
	fmt.Println(email)
	DB, err := database.ConnectToDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to the database",
		})
		return
	}

	var notes []models.Notes
	DB.Find(&notes, "email=?", email)

	if len(notes) == 0 {
		errorMessage := fmt.Sprintf("No Notes present for %s", name)
		c.JSON(http.StatusNotFound, gin.H{
			"message": errorMessage,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Name":  name,
			"Notes": notes,
		})
	}
}

func InsertNote(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	c.Bind(&body)
	_, email := GetEmail(c)
	notes := models.Notes{
		Email:       email,
		Heading:     email + body.Title,
		Title:       body.Title,
		Description: body.Description,
	}
	DB, _ := database.ConnectToDatabase()
	result := DB.Create(&notes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to insert note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note Inserted successfully",
	})

}

func GetNote(c *gin.Context) {
	title := c.Param("title")
	/*c.JSON(200,gin.H{
		"Title":title,
	})*/
	_, email := GetEmail(c)
	DB, _ := database.ConnectToDatabase()
	var notes models.Notes
	heading := email + title
	DB.Find(&notes, "heading=?", heading)

	if notes.ID == 0 {
		errorMessage := fmt.Sprintf("No Notes present for %s", title)
		c.JSON(http.StatusNotFound, gin.H{
			"message": errorMessage,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Title":       notes.Title,
			"Description": notes.Description,
		})
	}

}

func UpdateNote(c *gin.Context) {
	var body struct {
		Description string `json:"description"`
	}
	c.Bind(&body)
	title := c.Param("title")
	_, email := GetEmail(c)
	DB, _ := database.ConnectToDatabase()
	heading := email + title

	// Find the existing note
	var notes models.Notes
	result := DB.Where("heading = ?", heading).First(&notes)
	if result.Error != nil {
		errorMessage := fmt.Sprintf("No Notes present for %s", title)
		c.JSON(http.StatusNotFound, gin.H{
			"message": errorMessage,
		})
		return
	}

	// Update the description
	notes.Description = body.Description

	// Save the updated note
	result = DB.Save(&notes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to update note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      "Note updated successfully",
		"updated_note": notes, // This will include the updated note details in the response
	})
}
func DeleteNotes(c *gin.Context) {
	_, email := GetEmail(c)
	var notes models.Notes
	DB, _ := database.ConnectToDatabase()
	if err := DB.Where("email=?", email).Delete(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete records"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Records deleted successfully"})
}

func DeleteNote(c *gin.Context) {
	_, email := GetEmail(c)
	title := c.Param("title")
	heading := email + title

	DB, _ := database.ConnectToDatabase()

	var note models.Notes
	result := DB.Where("heading = ?", heading).First(&note)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch note"})
		return
	}

	if err := DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "note deleted successfully"})
}

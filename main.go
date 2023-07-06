package main

import (
	"fmt"
	"log"
	"milestones/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("milestones", getMilestones)
		v1.GET("milestone/:IdMilestone", getMilestoneById)
		v1.POST("milestone", addMilestone)
		v1.PUT("milestone/:IdMilestone", updateMilestone)
		v1.DELETE("milestone/:IdMilestone", deleteMilestone)
		v1.OPTIONS("milestone", options)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

func getMilestones(c *gin.Context) {

	milestones, err := models.GetMilestones(10)

	checkErr(err)

	if milestones == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": milestones})
	}
}

func getMilestoneById(c *gin.Context) {

	// grab the IdMilestone of the record want to retrieve
	IdMilestone := c.Param("IdMilestone")

	milestone, err := models.GetMilestoneById(IdMilestone)

	checkErr(err)
	// if the name is blank we can assume nothing is found
	if milestone.Milestone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": milestone})
	}
}

func addMilestone(c *gin.Context) {

	var json models.Milestone
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddMilestone(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func updateMilestone(c *gin.Context) {

	var json models.Milestone

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	IdMilestone, err := strconv.Atoi(c.Param("IdMilestone"))

	fmt.Printf("Updating IdMilestone %d", IdMilestone)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IdMilestone"})
	}

	success, err := models.UpdateMilestone(json, IdMilestone)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deleteMilestone(c *gin.Context) {

	IdMilestone := c.Param("IdMilestone")

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IdMilestone"})
	// }

	success, err := models.DeleteMilestone(IdMilestone)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func options(c *gin.Context) {

	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

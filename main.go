// PersonWeb
// By Jeremy Morgan
// https://www.allhandsontech.com/programming/golang/web-app-sqlite-go/
// See also: https://practicalgobook.net/posts/go-sqlite-no-cgo/

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
  "models"
  "log"
)

func main() {
  err := models.ConnectDatabase()
  checkErr(err)

  r := gin.Default()

  // API v1
  v1 := r.Group("/api/v1")
  {
    v1.GET("person", getPersons)
    v1.GET("person/:id", getPersonById)
    v1.POST("person", addPerson)
    v1.PUT("person/:id", updatePerson)
    v1.DELETE("person/:id", deletePerson)
    v1.OPTIONS("person", options)
  }

  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  r.Run()
}

func getPersons(c *gin.Context) {

	persons, err := models.GetPersons(10)
	checkErr(err)

	if persons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getPersonById " + id +" Called"})
}

func addPerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addPerson Called"})
}

func updatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updatePerson Called"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deletePerson " + id + " Called"})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options Called"})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
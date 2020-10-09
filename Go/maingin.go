package main

import (
	//"model"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  string
}

func main() {
	dsn := "docker:docker@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})

	//s := Person{Name: "Sean", Age: 50}
	//s.Name = "Sean"

	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HOLA GIN",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HOLA GIN",
		})
	})

	r.GET("/persons/:id", func(c *gin.Context) {
		id := c.Param("id")
		var d Person
		if err := db.Where("id = ?", id).First(&d).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Unscoped().Delete(&d)
	})
	//db.First(&d, id)

	//})

	r.GET("/persons", func(c *gin.Context) {
		var lis []Person
		db.Find(&lis)
		c.JSON(http.StatusOK, lis)
	})

	r.POST("/persons/", func(c *gin.Context) {
		d := Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
		db.Create(&d)

		c.JSON(200, gin.H{ // serializador de gin
			"name": d.Name,
			"age":  d.Age,
		})

	})

	// ELIMINAR PERSONA POR ID
	/*r.DELETE("/persons/:id", func(c *gin.Context) { //GET ALL
		id := c.Param("id")
		var d Person                 //Crear variable person
		db.First(&d, id)             //Encontrar persona
		db.Unscoped().Delete(&d, id) //Eliminar Persona
		c.JSON(http.StatusOK, d)
	})*/

	r.DELETE("/persons/:id", func(c *gin.Context) {
		id := c.Param("id")
		var d Person
		if err := db.Where("id = ?", id).First(&d).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Unscoped().Delete(&d)
	})

	r.PUT("/persons/:id", func(c *gin.Context) {
		id := c.Param("id")
		var d Person
		if err := db.First(&d, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		d.Name = c.PostForm("name")
		d.Age = c.PostForm("age")
		//c.BindJSON(&d)
		db.Save(&d)
		c.JSON(http.StatusOK, &d)
	})

	r.Run()

}

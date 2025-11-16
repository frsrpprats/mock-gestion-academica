package main

import (
	"net/http"
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
)

type Especialidad struct {
	Id           int    `fake:"{number:1}" json:"id"`
	Especialidad string `fake:"{jobtitle}" json:"especialidad"`
	Facultad     string `fake:"{company}" json:"facultad"`
	Universidad  string `fake:"{company}" json:"universidad"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	especialidades := generateMockData()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Mock API is running",
		})
	})

	r.GET("/api/v1/especialidades", func(c *gin.Context) {
		c.JSON(http.StatusOK, especialidades)
	})

	r.GET("/api/v1/especialidades/:id", func(c *gin.Context) {
		_, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}

		idx := gofakeit.Number(0, len(especialidades)-1)
		if idx != -1 {
			c.JSON(http.StatusOK, especialidades[idx])
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Especialidad not found"})

	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func generateMockData() []Especialidad {
	// Generar datos en idioma español
	gofakeit.Seed(0)
	especialidades := []Especialidad{}
	gofakeit.Slice(&especialidades)
	return especialidades
}

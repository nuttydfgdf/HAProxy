package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id         int
	Name       string `fake:"{firstname}"` // Any available function all lowercase
	Occupation string
}

func main() {
	/*
		// Hello world, the web server
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		}

		http.HandleFunc("/hello", helloHandler)
		log.Println("Listing for requests at http://localhost:8000/hello")
		log.Fatal(http.ListenAndServe(":8000", nil))
	*/

	users := []User{
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{"data": "hello world", "Env.HOST" : os.Getenv("HOST")})
	})

	// Routes
	//r.GET("/users", controllers.FindUsers)
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	r.Run(":9090")
}

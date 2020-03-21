package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default one shipped
	router := gin.Default()

	// Serve frontend static files
	//Point Golang to use the BUILD folder from React as the front-end
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Nice try buddy!",
			})

		})

	}
	// Our API will consit of just two routes

	//api.GET("/weather", weatherHandler)
	//This is the endpoint formatting for the time traveler on Dark Sky
	api.GET("/weather/:lat/:long/:time", weatherHandler)

	// Start and run the server
	//The port number has to come from the .env file based on the environment
	//Producion
	router.Run(":8080")
	//Dev
	//router.Run(":3000")
}

//Consume the weather API
func weatherHandler(c *gin.Context) {
	//start
	fmt.Println("Received params... ")
	//if lat, err := c.Param("lat"); err == nil {

	fmt.Println("lat: " + string(c.Param("lat")))
	fmt.Println("long: " + string(c.Param("long")))
	fmt.Println("time: " + string(c.Param("time")))
	//}
	//The API KEY value should be stored in the config file for a better maintainance
	APIKey := "b9d6d05e47ed429502486f5a8e56dc44/"
	var params = c.Param("lat") + "," + c.Param("long") + "," + c.Param("time")
	//Use ReGexp to replace all occurrence of the "&" in the params string
	re := regexp.MustCompile(`&`)
	param := re.ReplaceAllString(params, "")
	//The url for the Restfull service should be stored in the config file for a better maintainance
	url := "https://api.darksky.net/forecast/" + APIKey + param

	fmt.Println("URL: " + url)

	response, err := http.Get(url)
	//response, err := http.Get("https://api.darksky.net/forecast/b9d6d05e47ed429502486f5a8e56dc44/42.3601,-71.0589,255657600")
	if err != nil {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{
			//fmt.Printf("The HTTP request failed with error %s\n", err)
			"message": fmt.Errorf("decompress %v", err),
		})
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{
			//fmt.Fprint(w, data),
			"message": string(data),
		})
	}
	//finish
}

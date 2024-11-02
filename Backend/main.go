package main

import(
	"fmt"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"io"
	"time"

)

func getPrice(c *gin.Context) {
	id := c.Query("id")
	amount := c.Query("amount")	


	coingeckoUrl := fmt.Sprintf("https://pro-api.coinmarketcap.com/v2/tools/price-conversion?amount=%v&symbol=%v", amount, id)

	// Creating Client
	client := &http.Client{}
	req, err := http.NewRequest("GET", coingeckoUrl, nil)
    if err != nil {
        log.Println("Error creating request:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
        return
    }

	req.Header.Set("accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", "6785dae0-379a-4ef2-a3ea-998bb9a03372")

	// Send request	
	resp, err := client.Do(req)
    if err != nil {
        log.Println("Error making API call:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
    }
    defer resp.Body.Close()

	// Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
        return
    }

    // Return the response from the external API as JSON
    c.Header("Content-Type", "application/json")
    c.String(http.StatusOK, string(body))
}

func tokensList(c *gin.Context){
	url := "https://api.coingecko.com/api/v3/simple/supported_vs_currencies"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
        return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("x-cg-pro-api-key", "CG-BRausksF9ekuD2GJondFLLAm")

	// Send request	
	resp, err := client.Do(req)
    if err != nil {
        log.Println("Error making API call:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
    }
    defer resp.Body.Close()

	// Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error reading response body:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
        return
    }

    // Return the response from the external API as JSON
    c.Header("Content-Type", "application/json")
    c.String(http.StatusOK, string(body))

}

func main() {
	router := gin.Default()

	  // Configure CORS middleware
	  config := cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }

    // Apply the CORS middleware
    router.Use(cors.New(config))
	router.GET("/supportedList", tokensList)
	router.GET("/tokenprice", getPrice)
	router.Run("localhost:8080")
}
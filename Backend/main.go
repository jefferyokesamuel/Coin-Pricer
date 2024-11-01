package main

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"io"
	
)

func getPrice(c *gin.Context) {
	id := c.Query("id")	
	coingeckoUrl := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%v&vs_currencies=usd", id)

	// Creating Client
	client := &http.Client{}
	req, err := http.NewRequest("GET", coingeckoUrl, nil)
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
	router.GET("/tokenprice", getPrice)
	router.Run("localhost:8080")
}
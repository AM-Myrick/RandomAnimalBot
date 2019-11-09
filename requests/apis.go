package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetDogPic returns URL of a dog picture
func GetDogPic() string {
	var dogPic string
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal("Failed to get dog picture.")
	}

	defer resp.Body.Close()
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if result["status"] == "success" {
		if dogPic, ok := result["message"].(string); ok {
			fmt.Println(dogPic)
			return dogPic
		}
	}
	return dogPic
}

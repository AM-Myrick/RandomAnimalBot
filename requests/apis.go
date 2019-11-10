package requests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchDogPic returns URL of a dog picture
func FetchDogPic() string {
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
			return dogPic
		}
	}
	return dogPic
}

// FetchCatPic returns URL of a cat picture
func FetchCatPic() string {
	var catPic string
	resp, err := http.Get("https://aws.random.cat/meow")
	if err != nil {
		log.Fatal("Failed to get cat picture.")
	}

	defer resp.Body.Close()
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if catPic, ok := result["file"].(string); ok {
		return catPic
	}
	return catPic
}

// FetchFoxPic returns URL of a fox picture
func FetchFoxPic() string {
	var foxPic string
	resp, err := http.Get("https://randomfox.ca/floof/")
	if err != nil {
		log.Fatal("Failed to get fox picture.")
	}

	defer resp.Body.Close()
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	if foxPic, ok := result["image"].(string); ok {
		return foxPic
	}
	return foxPic
}

// FetchShibaPic returns URL of a shiba inu picture
func FetchShibaPic() string {
	var shibaPic string
	resp, err := http.Get("http://shibe.online/api/shibes")
	if err != nil {
		log.Fatal("Failed to get shiba picture.")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Failed to read response body.")
	}

	for _, letter := range body {
		// checking for the existence of quotes or brackets
		if letter == 34 || letter == 91 || letter == 93 || letter == 79 {
			continue
		}
		shibaPic += string(letter)
	}

	return shibaPic
}

package pterodactyl

import (
	"discord_pterodactyl_connector/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetStatus(config *config.Config) string {
	url := fmt.Sprintf("%sservers/%s/resources", config.PterodactylURL, config.ServerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "Error fetching server status"
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "Error fetching server status"
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return "Error fetching server status"
	}
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return "Error fetching server status"
	}
	status := result["attributes"].(map[string]interface{})["current_state"].(string)
	return fmt.Sprintf("Server status: %s", status)
}

package pterodactyl

import (
	"bytes"
	"discord_pterodactyl_connector/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func KillServer(config *config.Config) {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": "kill"}
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Error response from server: %s", body)
		return
	}
	log.Println("Server killed successfully")
}

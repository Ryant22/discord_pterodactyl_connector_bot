package pterodactyl

import (
	"discord_pterodactyl_connector/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStatus(config *config.Config) (string, error) {
	url := fmt.Sprintf("%sservers/%s/resources", config.PterodactylURL, config.ServerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error response from server: %s", body)
	}
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}
	status, ok := result["attributes"].(map[string]interface{})["current_state"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format: missing current_state")
	}
	return status, nil
}

package pterodactyl

import (
	"bytes"
	"discord_pterodactyl_connector/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendPowerSignal(signal string, config *config.Config) error {
	url := fmt.Sprintf("%sservers/%s/power", config.PterodactylURL, config.ServerID)
	payload := map[string]string{"signal": signal}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error response from server: %s", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server error: %s", body)
	}
	return nil
}

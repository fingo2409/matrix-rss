package matrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gomarkdown/markdown"
)

func SendMatrixMessage(server, roomID, token, message string) error {
	htmlMessage := markdown.ToHTML([]byte(message), nil, nil)

	url := fmt.Sprintf("%s/_matrix/client/r0/rooms/%s/send/m.room.message?access_token=%s", server, roomID, token)
	payload := map[string]interface{}{
		"msgtype":        "m.text",
		"body":           message,
		"format":         "org.matrix.custom.html",
		"formatted_body": string(htmlMessage),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal matrix payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "matrix-rss-bot/1.0")

	client := &http.Client{
		Timeout: 15 * time.Second, // optional, schöneres Fehlverhalten
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("matrix request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message: %s, body: %s", resp.Status, string(body))
	}
	return nil
}

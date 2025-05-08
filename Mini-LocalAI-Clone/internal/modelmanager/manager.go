package modelmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Manager struct {
	client  *http.Client
	baseURL string
}

func NewManager() *Manager {
	return &Manager{
		client:  &http.Client{},
		baseURL: "http://localhost:11434",
	}
}

func (m *Manager) ListModels() ([]ModelInfo, error) {
	resp, err := m.client.Get(m.baseURL + "/api/tags")
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// For debugging
	fmt.Printf("Raw API response: %s\n", string(body))

	var ollamaResp OllamaTagResponse
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert Ollama response to our ModelInfo format
	var models []ModelInfo
	for _, model := range ollamaResp.Models {
		// Check if the model supports multimodal based on family
		isMultiModal := false
		for _, family := range model.Details.Families {
			if family == "clip" || family == "llava" {
				isMultiModal = true
				break
			}
		}

		models = append(models, ModelInfo{
			Name:         model.Name,
			IsMultiModal: isMultiModal,
		})
	}

	return models, nil
}

func (m *Manager) Run(actions []Action) (*ActionResult, error) {
	models, err := m.ListModels()
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return nil, fmt.Errorf("no models available. Please pull a model using 'ollama pull <model-name>'")
	}

	// Return the selected model info
	return &ActionResult{
		ModelName:    models[0].Name,
		IsMultiModal: models[0].IsMultiModal,
	}, nil
}

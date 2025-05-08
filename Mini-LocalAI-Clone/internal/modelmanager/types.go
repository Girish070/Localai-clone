package modelmanager

type ModelInfo struct {
	Name         string `json:"name"`
	IsMultiModal bool   `json:"is_multi_modal"`
}

type ActionResult struct {
	ModelName    string
	IsMultiModal bool
	Error        error
}

type Action string

const (
	ActionInstall Action = "install"
	ActionManage  Action = "manage"
	ActionMonitor Action = "monitor"
	ActionUpdate  Action = "update"
	ActionDelete  Action = "delete"
	ActionChat    Action = "chat"
)

// Ollama API response types
type OllamaTagResponse struct {
	Models []struct {
		Name    string `json:"name"`
		Details struct {
			Format             string   `json:"format"`
			Family             string   `json:"family"`
			Families           []string `json:"families"`
			Parameter_size     string   `json:"parameter_size"`
			Quantization_level string   `json:"quantization_level"`
		} `json:"details"`
	} `json:"models"`
}

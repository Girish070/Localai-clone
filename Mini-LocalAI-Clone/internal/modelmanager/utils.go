package modelmanager

import (
	"fmt"
)

func PrintActionResult(result *ActionResult, err error) error {
	if err != nil {
		return fmt.Errorf("action failed: %w", err)
	}

	if result == nil {
		return fmt.Errorf("no result available")
	}

	fmt.Printf("Selected model: %s (MultiModal: %v)\n", result.ModelName, result.IsMultiModal)
	return nil
}

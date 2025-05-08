package main

import (
	"github.com/AkhilSharma90/GO-Native-LLM/internal/modelmanager"
)

func (cfg *gollamaConfig) ollamanager() {
	manager := modelmanager.NewManager()

	var actions []modelmanager.Action
	if cfg.Install {
		actions = append(actions, modelmanager.ActionInstall)
	}
	if cfg.Manage {
		actions = append(actions, modelmanager.ActionManage)
	}
	if cfg.Monitor {
		actions = append(actions, modelmanager.ActionMonitor)
	}

	result, err := manager.Run(actions)
	err = modelmanager.PrintActionResult(result, err)
}

package chat

import (
	"fmt"
	"strings"
	"time"

	"github.com/AkhilSharma90/GO-Native-LLM/internal/client"
	"github.com/AkhilSharma90/GO-Native-LLM/internal/modelmanager"
	"github.com/AkhilSharma90/GO-Native-LLM/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func NewChatSettingsForm() (client.Chat, error) {
	var newChatSettings client.Chat

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Description("Chat Title").
				Placeholder("Chat Title").
				Validate(func(title string) error {
					trimmedTitle := strings.TrimSpace(title)
					if trimmedTitle == "" {
						return fmt.Errorf("Chat Title cannot be empty")
					}
					return nil
				}).
				Value(&newChatSettings.ChatTitle),
			huh.NewText().
				Title("System Message").
				Placeholder("(Optional) Leave empty if you don't want to set a system message.").
				Value(&newChatSettings.SystemMessage),
			huh.NewConfirm().
				Title("Anonymous Chat").
				Description("Do you want to create an anonymous chat? (Messages will not be saved)").
				Value(&newChatSettings.IsAnonymous),
		),
	).WithProgramOptions(tea.WithAltScreen())

	err := form.Run()
	if err != nil {
		return client.Chat{}, fmt.Errorf("error: %w", err)
	}

	manager := modelmanager.NewManager()
	result, err := manager.Run([]modelmanager.Action{modelmanager.ActionChat})
	if err != nil {
		utils.PrintError(err, true)
	}

	newChatSettings.ModelName = result.ModelName
	newChatSettings.IsMultiModal = result.IsMultiModal

	if !newChatSettings.IsAnonymous {
		newChatSettings.ID = GenerateChatID()

		if err := client.GollamaInstance.CreateChat(newChatSettings); err != nil {
			return client.Chat{}, fmt.Errorf("error creating chat: %w", err)
		}
	}

	newChatSettings.UpdatedAt = time.Now()
	return newChatSettings, nil
}

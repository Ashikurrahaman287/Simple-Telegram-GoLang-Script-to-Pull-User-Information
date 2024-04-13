package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Replace this with your own Telegram API token
	apiToken := "https://ashikurrahaman.com/"

	// Create a new Telegram bot API client
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Fatalf("Failed to create bot API: %v", err)
	}

	// Connect to Telegram
	bot.Debug = true // Enable debugging to log HTTP requests
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Collect usernames from a Telegram group
	usernames, err := collectUsernames(bot, "@SpudblockBlockchainCommunity")
	if err != nil {
		log.Fatalf("Failed to collect usernames: %v", err)
	}

	// Save usernames to a file
	if err := saveUsernamesToFile("spudblocks_Pull_Script.txt", usernames); err != nil {
		log.Fatalf("Failed to save usernames to file: %v", err)
	}

	log.Println("Usernames saved to spudblocks_Pull_Script.txt successfully.")
}

// collectUsernames collects usernames from a Telegram group
func collectUsernames(bot *tgbotapi.BotAPI, groupName string) ([]string, error) {
	var usernames []string

	// Get the members of the group
	members, err := bot.GetChatMembersConfig(groupName)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat members: %v", err)
	}

	// Extract usernames
	for _, member := range members {
		if member.User != nil && member.User.UserName != "" {
			usernames = append(usernames, member.User.UserName)
		}
	}

	return usernames, nil
}

// saveUsernamesToFile saves usernames to a file
func saveUsernamesToFile(filename string, usernames []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	for _, username := range usernames {
		_, err := file.WriteString(username + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	return nil
}

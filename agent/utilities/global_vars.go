package utilities

import (
	"errors"
	bot "kotosBidAgent"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define global variables
var (
	// Agent specific
	AGENT_NAME      = ""
	AGENT_ATTRIBUTE = ""
	AGENT_ADMIN_ID  = ""
	ADMIN_LIST      []string
	Cwtchbot        *bot.CwtchBot

	// Auction community membership
	AuctionName           = ""
	AuctionCommunityID    = 0
	AuctionCommunityOnion = ""
	AUCTION_MGR_URI       = ""

	EntityOnion = ""
)

func SetGlobalVars() error {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	// Set agents vars
	AGENT_NAME = os.Getenv("NAME")
	if AGENT_NAME == "" {
		log.Println("Error: AGENT_NAME is empty")
		return errors.New("AGENT_NAME is empty")
	}

	AGENT_ATTRIBUTE = os.Getenv("ATTRIBUTE")
	if AGENT_ATTRIBUTE == "" {
		log.Println("Error: AGENT_ATTRIBUTE is empty")
		return errors.New("AGENT_ATTRIBUTE is empty")
	}

	AGENT_ADMIN_ID = os.Getenv("ADMIN")
	if AGENT_ADMIN_ID == "" {
		log.Println("Error: AGENT_ADMIN_ID is empty")
		return errors.New("AGENT_ADMIN_ID is empty")
	}

	// Set bot_admin_list
	ADMIN_LIST = append(ADMIN_LIST, AGENT_ADMIN_ID)

	// Auction community
	AuctionName = os.Getenv("AUCTION_NAME")
	if AuctionName == "" {
		log.Println("Error: AUCTION_NAME is empty")
		return errors.New("AUCTION_NAME is empty")
	}

	// Auction Data Manager
	AUCTION_MGR_URI = os.Getenv("AUCTION_MGR_URI")
	if AUCTION_MGR_URI == "" {
		log.Println("Error: AUCTION_MGR_URI is empty")
		return errors.New("AUCTION_MGR_URI is empty")
	}

	return nil
}

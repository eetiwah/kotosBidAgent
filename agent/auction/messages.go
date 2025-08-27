package auction

import (
	"encoding/json"
	"fmt"
	"kotosBidAgent/agent/group"
	"log"
	"strings"
)

func Messages(data string, conversationID int, onion string) {
	var groupMsg group.GroupMessage

	// convert string to []byte
	jsonData := []byte(data)

	// Unmarshal string to JSON cmd structure
	if err := json.Unmarshal(jsonData, &groupMsg); err != nil {
		msg := fmt.Sprintf("Error: unmarshalling failure = %v\n", err)
		log.Println(msg)
		return
	}

	// Marshal Data back to JSON and unmarshal into the correct struct
	dataBytes, err := json.Marshal(groupMsg.Data)
	if err != nil {
		msg := fmt.Sprintf("Error: failed to marshal data: %v\n", err)
		log.Println(msg)
		return
	}

	// Determine auction message type
	switch strings.ToLower(groupMsg.Type) {

	case "create_auction":
		CreateAuctionObj(dataBytes)

	case "start_auction":
		StartAuction(dataBytes)

	case "stop_auction":
		StopAuction(dataBytes)

	default:
		log.Printf("Auction MessageType error: %v from %d", groupMsg.Type, conversationID)
	}
}

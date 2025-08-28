package auction

import (
	"encoding/json"
	"log"
)

// Not if this is going to work having two definitions, but they are the same
type GroupMessage struct {
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

// UnmarshalJSON customizes JSON unmarshaling for GroupMessage
func (g *GroupMessage) UnmarshalJSON(data []byte) error {
	// Temporary struct to capture type and version
	type Alias struct {
		Type    string          `json:"type"`
		Version string          `json:"version"`
		Data    json.RawMessage `json:"data"` // Use RawMessage to delay Data parsing
	}
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		log.Printf("Error unmarshaling into Alias: %v", err)
		return err
	}

	// Assign type and version
	g.Type = alias.Type
	g.Version = alias.Version

	// Parse Data based on Type
	switch g.Type {
	case "ping_auction":
		g.Data = ""

	case "create_auction":
		var o_req AuctionObject
		if err := json.Unmarshal(alias.Data, &o_req); err != nil {
			log.Printf("Error unmarshaling Data into AuctionObject: %v", err)
			return err
		}
		g.Data = o_req

	default:
		log.Printf("Unsupported message type: %s, storing raw data", g.Type)
		g.Data = alias.Data
	}
	return nil
}

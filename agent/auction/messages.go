package auction

import (
	"log"
	"strings"
)

func Messages(data string, conversationID int, onion string) {
	cmdList := strings.Split(data, " ")
	/*
		var groupMsg group.GroupMessage

		// convert string to []byte
		//jsonData := []byte(data)

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
	*/

	// Determine auction message type
	//switch strings.ToLower(groupMsg.Type) {
	switch strings.ToLower(cmdList[0]) {

	case "ping_auction":
		log.Println("ping_auction received")

	case "create_auction":
		log.Println("create_auction received")
		//CreateAuctionObj(dataBytes)
		//CreateAuctionObj([]byte(cmdList[1]))

	case "start_auction":
		log.Println("start_auction received")
		//StartAuction(dataBytes)
		//StartAuction([]byte(cmdList[1]))

	case "stop_auction":
		log.Println("stop_auction received")
		//StopAuction(dataBytes)
		//StopAuction([]byte(cmdList[1]))

	default:
		//log.Printf("Auction MessageType error: %v from %d", groupMsg.Type, conversationID)
		log.Printf("Auction MessageType error: %v from %d", cmdList[0], conversationID)
	}
}

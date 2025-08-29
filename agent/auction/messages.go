package auction

import (
	"encoding/json"
	"fmt"
	"kotosBidAgent/agent/group"
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
		// log.Printf("data received = %s", cmdList[1])
		CreateAuctionObj([]byte(cmdList[1]))

	case "start_auction":
		log.Println("start_auction received")
		// log.Printf("data received = %s", cmdList[1])
		//StartAuction(dataBytes)
		err := StartAuction(cmdList[1])
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		// Generate the bid
		bid, err := GenerateBid(cmdList[1])
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		log.Println("Bid was generated")

		// Add bid to bid store
		err = AddBid(bid)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		log.Println("Bid was stored")

		byteData, err := json.Marshal(bid)
		if err != nil {
			errMsg := fmt.Sprintf("AddBid: status code %d", err)
			log.Println(errMsg)
		}

		// Rebuild command
		command := fmt.Sprintf("%s %s", "bid_offer", string(byteData))

		// Send the command to the community
		err = group.SendMessage([]byte(command))
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		log.Println("Bid was sent")

	case "stop_auction":
		log.Println("stop_auction received")
		// log.Printf("data received = %s", cmdList[1])
		//StopAuction(dataBytes)
		err := StopAuction(cmdList[1])
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println("Auction was stopped")

	case "set_auction_winner":
		log.Println("set_auction_winner received")
		err := SetAuctionWinner(cmdList[1], cmdList[2])
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Printf("Auction %s winner was set %s", cmdList[1], cmdList[2])

	default:
		//log.Printf("Auction MessageType error: %v from %d", groupMsg.Type, conversationID)
		log.Printf("Auction MessageType error: %v from %d", cmdList[0], conversationID)
	}
}

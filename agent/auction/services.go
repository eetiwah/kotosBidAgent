package auction

import (
	"encoding/json"
	"errors"
	"fmt"
	"kotosBidAgent/agent/group"
	"log"
)

// Auction management messages from admin

func Get(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_auction auctionId"
		}

		obj, err := GetAuctionObj(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return fmt.Sprintf("Error: get_auction json marshal: %v", err)
		}
		return fmt.Sprintf("Auction: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func List(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_auction_list"
		}

		objList, err := GetAuctionList()
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(objList)
		if err != nil {
			return fmt.Sprintf("Error: get_auction_list json marshal: %v", err)
		}
		return fmt.Sprintf("Auction List: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func GetBid(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_bid bidId"
		}

		obj, err := GetBidObj(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return fmt.Sprintf("Error: get_bid json marshal: %v", err)
		}
		return fmt.Sprintf("Bid: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

func BidList(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing "

	case 2:
		if commandList[1] == "-help" {
			return "usage: get_bid_list auctionId"
		}

		objList, err := GetBidList(commandList[1])
		if err != nil {
			return err.Error()
		}

		jsonStr, err := json.Marshal(objList)
		if err != nil {
			return fmt.Sprintf("Error: get_bid_list json marshal: %v", err)
		}
		return fmt.Sprintf("Bid List: %s", jsonStr)

	default:
		return "Error: parameter mismatch"
	}
}

// Not using this one so far

func Bid_Event_Sent(bidObj BidObject) error {
	// Create a group message of the "bid_send" type
	groupMsg := group.GroupMessage{
		Type:    "bid_response",
		Version: "1.0",
		Data:    bidObj,
	}

	// Serialize the message
	dataBytes, err := json.Marshal(groupMsg)
	if err != nil {
		errMsg := fmt.Sprintf("Bid_Event_Sent: marshalling group message failed: %v", err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	err = group.SendMessage(dataBytes)
	if err != nil {
		errMsg := fmt.Sprintf("Bid_Event_Sent: send message error: %v", err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	log.Printf("Bid_Event_Sent: Bid %s was submitted for auction %s\n", bidObj.BidId, bidObj.AuctionId)
	return nil

	/*
	   // Get the conversation
	   conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)

	   	if err != nil {
	   		errMsg := fmt.Sprintf("Bid_Event_Sent: failed to find conversation for: %s, err: %v", utilities.AuctionCommunityOnion, err)
	   		log.Println(errMsg)
	   		//sendErrorMessage(conversation.ID, "Thread_Services: error", errMsg)
	   		return errors.New(errMsg)
	   	}

	   // Send response to the group/community
	   _, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(dataBytes))))

	   	if err != nil {
	   		log.Printf("Bid_Event_Sent: send message error: %v", err)
	   	}

	   log.Printf("Bid_Event_Sent: Bid %s was submitted for auction %s\n", bidObj.BidId, bidObj.AuctionId)
	   return nil
	*/
}

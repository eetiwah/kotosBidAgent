package auction

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
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

func List() string {
	objList, err := GetAuctionList()
	if err != nil {
		return err.Error()
	}

	jsonStr, err := json.Marshal(objList)
	if err != nil {
		return fmt.Sprintf("Error: get_auction_list json marshal: %v", err)
	}
	return fmt.Sprintf("Auction List: %s", jsonStr)
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

func GenerateBid(auctionId string) (BidObject, error) {
	bidObject := BidObject{
		BidId:     uuid.NewString(),
		AuctionId: auctionId,
		Price:     "5.00",
	}

	/*
		groupMsg := GroupMessage{
			Type:    "bid_offer",
			Version: "1.0",
			Data:    bidObject,
		}

		dataBytes, err := json.Marshal(groupMsg)
		if err != nil {
			return dataBytes, err
		}

	*/

	return bidObject, nil
}

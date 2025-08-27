package auction

// Auction management messages from admin

// Auction messages from community
/*
func Add_Order_Event_Received(conversationID int, onion string, byteData []byte) {
	err := AddOrder(byteData)
	if err != nil {
		log.Printf("Add_Order_Event_Received: %s\n", err.Error())
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return
	}

	var orderObj OrderObject
	if err := json.Unmarshal(byteData, &orderObj); err != nil {
		log.Printf("Add_Start_Event_Received: failed to unmarshal addObj: %v\n", err)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	log.Printf("Order %s was added to Order Collection\n", orderObj.OrderId)
}
*/

/*
func Auction_Start_Event_Received(conversationID int, onion string, byteData []byte) {
	err := AddAuction(byteData)
	if err != nil {
		log.Printf("Auction_Start_Event_Received: %s\n", err.Error())
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return
	}

	var auctionObj AuctionObject
	if err := json.Unmarshal(byteData, &auctionObj); err != nil {
		log.Printf("Auction_Start_Event_Received: failed to unmarshal addObj: %v\n", err)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	log.Printf("Auction %s was added to Auction Collection\n", auctionObj.Id)
}

func Auction_Stop_Event_Received(conversationID int, onion string, byteData []byte) {
	err := StopAuction(byteData)
	if err != nil {
		log.Printf("Auction_Stop_Event_Received: %s\n", err.Error())
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return
	}

	var auctionEnd AuctionEnd
	if err := json.Unmarshal(byteData, &auctionEnd); err != nil {
		log.Printf("Auction_Stop_Event_Received: failed to unmarshal addObj: %v\n", err)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	log.Printf("Auction %s was stopped\n", auctionEnd.AuctionID)
}

func Auction_Winner_Event_Received(conversationID int, onion string, byteData []byte) {
	err := SetAuctionWinner(byteData)
	if err != nil {
		log.Printf("Auction_Winner_Event_Received: %s\n", err.Error())
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return
	}

	var auctionWinner AuctionWinner
	if err := json.Unmarshal(byteData, &auctionWinner); err != nil {
		log.Printf("Auction_Winner_Event_Received: failed to unmarshal addObj: %v\n", err)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	log.Printf("The winner for Auction: %s was bid: %s\n", auctionWinner.AuctionID, auctionWinner.BidID)

	// Is this bidId ours? If so, we won
	if err := OurBidId(auctionWinner.BidID); err != nil {
		log.Printf("Auction_Winner_Event_Received: we did not win this auction: %s\n", auctionWinner.AuctionID)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}

	// Yes, we won, add jobObject to jobCollection
	if err := AddJob(auctionWinner.AuctionID); err != nil {
		log.Printf("Auction_Winner_Event_Received: failed to add new job for auction: %s\n", auctionWinner.AuctionID)
		//sendErrorMessage(conversationID, "Thread_Services: add_received error", msg)
		return
	}
}

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
		log.Printf(errMsg)
		return errors.New(errMsg)
	}

	log.Printf("Bid_Event_Sent: Bid %s was submitted for auction %s\n", bidObj.BidId, bidObj.AuctionId)
	return nil

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
}

*/

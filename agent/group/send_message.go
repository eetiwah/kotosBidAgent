package group

import (
	"errors"
	"fmt"
	"kotosBidAgent/agent/utilities"
	"log"

	"cwtch.im/cwtch/model"
)

func SendMessage(data []byte) error {
	// Get the conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AuctionCommunityOnion)
	if err != nil {
		errMsg := fmt.Sprintf("Error: sendmessage failed to find conversation for: %s, err: %v", utilities.AuctionCommunityOnion, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send response to the group/community
	_, err = utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, string(data))))
	return err
}

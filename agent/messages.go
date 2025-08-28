package main

import (
	"fmt"
	"kotosBidAgent/agent/admin"
	"kotosBidAgent/agent/auction"
	"kotosBidAgent/agent/utilities"
	"strings"

	"cwtch.im/cwtch/model"
)

func Messages(data string) string {
	cmd := strings.Split(data, " ")
	fmt.Printf("Command received: %s\n", cmd[0])

	switch strings.ToLower(cmd[0]) {

	// *** Connectivity check *** //

	case "ping":
		result := admin.Ping()
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	// *** Admin operations *** //

	case "add_admin":
		result := admin.AddAdmin(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "get_admin_list":
		result := admin.GetAdminList(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "remove_admin":
		result := admin.RemoveAdmin(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	// *** Contact Operations *** //

	case "add_contact":
		result := admin.AddContact(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "get_contact_list":
		result := admin.GetContactList()
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "contact_status":
		result := admin.GetContactStatus(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

		// *** Auction Operations *** //

	case "get_auction":
		result := auction.Get(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "get_auction_list":
		result := auction.List(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "get_bid":
		result := auction.GetBid(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	case "get_bid_list":
		result := auction.BidList(cmd)
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, result))

	default:
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, "Error: unrecognized command"))

	}
}

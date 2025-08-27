package admin

import (
	"encoding/json"
	"fmt"
	"kotosBidAgent/agent/utilities"
	"log"

	"cwtch.im/cwtch/model"
	"cwtch.im/cwtch/protocol/connections"
)

func Ping() string {
	return "pong"
}

func AddAdmin(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing adminID"

	case 2:
		if commandList[1] == "-help" {
			return "usage: addadmin adminID"
		} else {
			newadmin := commandList[1]
			utilities.ADMIN_LIST = append(utilities.ADMIN_LIST, newadmin)
			return "Admin was added"
		}

	default:
		return "Error: parameter mismatch"
	}
}

func GetAdminList(commandList []string) string {
	switch len(commandList) {
	case 1:
		return getList(utilities.ADMIN_LIST)

	case 2:
		if commandList[1] == "-help" {
			return "usage: getadminlist"
		} else {
			return "Error: parameter mismatch"
		}

	default:
		return "Error: parameter mismatch"
	}
}

func RemoveAdmin(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing adminNtKID"

	case 2:
		if commandList[1] == "-help" {
			return "usage: removeadmin admin_ntkID"
		} else {
			admin := commandList[1]
			utilities.ADMIN_LIST = removeFromList(utilities.ADMIN_LIST, admin)
			return "Admin was removed"
		}

	default:
		return "Error: parameter mismatch"
	}
}

func AddContact(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing peerID"

	case 2:
		if commandList[1] == "-help" {
			return "usage: addcontact peerID"
		}

		id := commandList[1]
		err := utilities.Cwtchbot.Peer.ImportBundle(id)
		if err != nil {
			// Check if the error message indicates success
			if err.Error() != "importBundle.success" {
				return fmt.Sprintf("importing of contact failed: %s, error: %v", id, err)
			}

			// If the error message is "importBundle.success", treat it as a success
			return fmt.Sprintf("Contact was created: %s", id)

		} else {
			return "Error: this process path should have been used to add contact"
		}

	default:
		return "Error: parameter mismatch"
	}
}

func GetContactList() string {
	conversations, err := utilities.Cwtchbot.Peer.FetchConversations()
	if err != nil {
		log.Printf("Error: failed to retrieve contact list: %v", err)
		return "Error: was not able to retrieve contact list"
	}

	if len(conversations) <= 0 {
		log.Println("Warning: contact list is empty")
		return "Error: contact list empty"
	}

	var nameList []string
	for _, conversation := range conversations {
		if conversation.Handle != "" {
			nameList = append(nameList, conversation.Handle)
		}
	}

	if len(nameList) == 0 {
		log.Println("Warning: no valid contact names found in conversations")
		return "Error: no valid contact names found"
	}

	jsonData, err := json.Marshal(nameList)
	if err != nil {
		log.Printf("Error: marshalling nameList failed: %v", err)
		return "Error: marshalling nameList failed"
	}

	return string(jsonData)
}

func GetContactStatus(commandList []string) string {
	switch len(commandList) {
	case 1:
		return "Error: missing contactID"

	case 2:
		if commandList[1] == "-help" {
			return "usage: getcontactstatus peerID"
		} else {
			connectionState := utilities.Cwtchbot.Peer.GetPeerState(commandList[1])
			switch connectionState {
			case connections.DISCONNECTED:
				return "Peer is not online"

			case connections.AUTHENTICATED:
				return "Peer is online"

			case connections.CONNECTED:
				return "Peer is online"

			case connections.CONNECTING:
				return "Peer is connecting"

			default:
				return "Peer state is unknown"
			}
		}

	default:
		return "Error: parameter mismatch"
	}
}

func SendAdminMsg(msg string) {
	// Check peer online
	connectionState := utilities.Cwtchbot.Peer.GetPeerState(utilities.AGENT_ADMIN_ID)
	if connectionState == connections.DISCONNECTED {
		log.Println("Error: peer is not online")
		return
	}

	// Get conversation
	conversation, err := utilities.Cwtchbot.Peer.FetchConversationInfo(utilities.AGENT_ADMIN_ID)
	if err != nil {
		log.Printf("Error: admin is not a contact: " + utilities.AGENT_ADMIN_ID)
		return
	}

	// Send message
	packMsg := string(utilities.Cwtchbot.PackMessage(model.OverlayChat, msg))
	if _, err := utilities.Cwtchbot.Peer.SendMessage(conversation.ID, packMsg); err != nil {
		log.Printf("Error: cannot send message to admin: %v", err)
		return
	}
	log.Printf("Sent message to adminL %s", msg)
}

func getList(list []string) string {
	if len(list) > 0 {
		jsonBytes, err := json.Marshal(list)
		if err != nil {
			return "Error: encoding of list"
		} else {
			return string(jsonBytes)
		}
	} else {
		return "The list was empty"
	}
}

func removeFromList(usersList []string, userToRemove string) []string {
	for i, user := range usersList {
		if user == userToRemove {
			// Remove the user from the slice by slicing it to exclude the user
			return append(usersList[:i], usersList[i+1:]...)
		}
	}
	// If userToRemove is not found, return the original slice
	return usersList
}

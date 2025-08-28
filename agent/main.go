package main

import (
	"encoding/json"
	"errors"
	"fmt"
	bot "kotosBidAgent"
	"kotosBidAgent/agent/admin"
	"kotosBidAgent/agent/auction"
	"kotosBidAgent/agent/group"
	"kotosBidAgent/agent/utilities"
	"log"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"syscall"

	"cwtch.im/cwtch/event"
	"cwtch.im/cwtch/model"
	"cwtch.im/cwtch/model/attr"
	"cwtch.im/cwtch/model/constants"

	_ "github.com/mutecomm/go-sqlcipher/v4"
)

const (
	TextMessageOverlay       = 1
	ActionableMessageOverlay = 5
	SuggestContactOverlay    = 100
	InviteGroupOverlay       = 101
)

type OverlayEnvelope struct {
	ConversationID int
	Onion          string
	Overlay        int    `json:"o"`
	Data           string `json:"d"`
}

// Message holds parts data
type Message struct {
	O int    `json:"o"`
	D string `json:"d"`
}

// Instantiate new agent
func instantiateAgent() error {
	botpath := "/" + utilities.AGENT_NAME + "/"

	switch runtime.GOOS {
	case "windows":
		utilities.Cwtchbot = bot.NewCwtchBot(path.Join("./tor/win", botpath), utilities.AGENT_NAME)

	case "linux":
		_path := path.Join("./tor/linux", botpath)
		utilities.Cwtchbot = bot.NewCwtchBot(_path, utilities.AGENT_NAME)

	default:
		return fmt.Errorf("operating system not support = %v", runtime.GOOS)
	}

	utilities.Cwtchbot.Launch() // Need some error check here

	// Set Some Profile Information
	utilities.Cwtchbot.Peer.SetScopedZonedAttribute(attr.PublicScope, attr.ProfileZone, constants.Name, utilities.AGENT_NAME)
	utilities.Cwtchbot.Peer.SetScopedZonedAttribute(attr.PublicScope, attr.ProfileZone, constants.ProfileAttribute1, utilities.AGENT_ATTRIBUTE)
	utilities.Cwtchbot.Peer.SetScopedZonedAttribute(attr.PublicScope, attr.ProfileZone, constants.ProfileAttribute2, utilities.AGENT_ATTRIBUTE)
	utilities.Cwtchbot.Peer.SetScopedZonedAttribute(attr.PublicScope, attr.ProfileZone, constants.ProfileAttribute3, utilities.AGENT_ATTRIBUTE)

	// Display address
	utilities.EntityOnion = utilities.Cwtchbot.Peer.GetOnion()
	log.Printf("%s address: %v\n", utilities.AGENT_NAME, utilities.EntityOnion)

	return nil
}

func isContact(id string) error {
	_, err := utilities.Cwtchbot.Peer.FetchConversationInfo(id)
	if err != nil {
		importErr := utilities.Cwtchbot.Peer.ImportBundle(id)
		if importErr != nil {
			// Check if the error message indicates success
			if importErr.Error() != "importBundle.success" {
				msg := fmt.Sprintf("importing of contact failed: %s, error: %v", id, importErr)
				return errors.New(msg)
			}
			// If the error message is "importBundle.success", treat it as a success
			log.Printf("Contact was created: %s", id)
			utilities.ADMIN_LIST = append(utilities.ADMIN_LIST, id)
			return nil
		}
		log.Printf("Contact was created: %s", id)
		return nil

	} else {
		log.Printf("Contact is already a contact: %s", id)
		return nil
	}
}

func main() {
	// Set global variables
	log.Println("Setting global vars...")
	err := utilities.SetGlobalVars()
	if err != nil {
		log.Printf("Error loading .env file: %s", err.Error())
		return
	}

	// Create the bot
	log.Println("Instantiating agent...")
	err = instantiateAgent()
	if err != nil {
		log.Printf("Error: instantiating the bot, %s", err.Error())
		return
	}

	// Is adminID a contact? If not add
	log.Println("Determining if the admin a contact...")
	err = isContact(utilities.AGENT_ADMIN_ID)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	// Handle graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Processing loop
	go func() {
		log.Println("Starting message queue")
		for {
			message := utilities.Cwtchbot.Queue.Next()

			switch message.EventType {

			// This does not occur with out group invite
			case event.InvitePeerToGroup:
				conversation, _ := utilities.Cwtchbot.Peer.FetchConversationInfo(message.Data[event.RemotePeer])
				log.Printf("Received InvitePeerToGroup from %v\n", conversation.Handle)

			case event.ContactCreated:
				conversation, _ := utilities.Cwtchbot.Peer.FetchConversationInfo(message.Data[event.RemotePeer])
				log.Printf("Received contact request from %v\n", conversation.Handle)

				err := utilities.Cwtchbot.Peer.AcceptConversation(conversation.ID)
				if err != nil {
					msg := fmt.Sprintf("Contact request failed from %v\n", conversation.Handle)
					if conversation.Handle != utilities.AGENT_ADMIN_ID { // Cannot send a msg to admin, if it sent the request
						admin.SendAdminMsg(msg)
					}
					log.Println(msg)
				} else {
					msg := fmt.Sprintf("Contact request was accepted from %v\n", conversation.Handle)
					if conversation.Handle != utilities.AGENT_ADMIN_ID { // Cannot send a msg to admin, if it sent the request
						admin.SendAdminMsg(msg)
					}
					log.Println(msg)
				}

			case event.NewMessageFromPeer:
				log.Println("NewMessageFromPeer")
				conversation, _ := utilities.Cwtchbot.Peer.FetchConversationInfo(message.Data[event.RemotePeer])
				envelope := Unwrap(conversation.ID, message.Data[event.RemotePeer], message.Data[event.Data])

				// Check if this is a response or not
				if envelope.Data != "Error:" && envelope.Data != "Success:" {
					switch envelope.Overlay {

					case TextMessageOverlay:
						reply := Messages(envelope.Data)
						utilities.Cwtchbot.Peer.SendMessage(conversation.ID, reply)

					case InviteGroupOverlay:
						reply := group.InviteGroup(envelope.Data)
						utilities.Cwtchbot.Peer.SendMessage(conversation.ID, reply)

					case SuggestContactOverlay:
						utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, "Received Suggest Contact Overlay request")))

					default:
						utilities.Cwtchbot.Peer.SendMessage(conversation.ID, string(utilities.Cwtchbot.PackMessage(model.OverlayChat, "Error: unrecognized command")))
					}
				} else {
					// The response will be logged
					log.Printf("Response: %s\n", envelope.Data)
				}

			case event.NewMessageFromGroup:
				// Convert ConversationID from string to int
				conversationID, err := strconv.Atoi(message.Data[event.ConversationID])
				if err != nil {
					log.Printf("error: failed to convert string to textto int: %v", err)
				}

				// Create an envelope object to hold the message information
				envelope := Unwrap(conversationID, message.Data[event.RemotePeer], message.Data[event.Data])

				// Is this an auction message?
				if conversationID == utilities.AuctionCommunityID {
					//log.Println("Auction message")

					switch envelope.Overlay {
					case TextMessageOverlay:
						//log.Println("*** We are not processing the prior auction messages...")
						auction.Messages(envelope.Data, envelope.ConversationID, envelope.Onion)

					default:
						log.Println("Error: unrecognized command")
					}
				} else {
					log.Println("Error: unrecoginized community message")
				}

			case event.PeerStateChange:
				log.Printf("PeerStateChange: %s\n", message.Data[event.ConnectionState])
				log.Printf("Remote Peer = %v\n", message.Data[event.RemotePeer])

			case event.ServerStateChange:
				log.Printf("ServerStateChange: %s\n", message.Data[event.ConnectionState])
				log.Printf("Group Server = %v\n", message.Data[event.GroupServer])

				// Has group server Synced occurred?
				if message.Data[event.ConnectionState] == "Synced" /*&& groupServer ==  message.Data[event.GroupServer]*/ {
					if _onion, _ID, err := getGroupHandle(utilities.AuctionName); err == nil {
						utilities.AuctionCommunityOnion = _onion
						utilities.AuctionCommunityID = _ID
						log.Printf("*** %s is now available\n", utilities.AuctionName)

						// Start the auction Watcher after the group server is operational
						//log.Println("Auction Watcher is not running, we are waiting for messages...")
						//go auction.Watcher()
						break
					}

					log.Printf("No conversation found with local.profile.name: %s\n", err.Error())
				}

			case event.PeerAcknowledgement:
				log.Println("PeerAcknowledgement")
				log.Printf("Remote Peer = %v\n", message.Data[event.RemotePeer])

			case event.SendMessageToPeerError:
				log.Println("SendMessageToPeerError occurred")

			case event.SendRetValMessageToPeer:
				log.Println("SendRetValMessageToPeer received")
				// We need to dig into this, but it does not effect the functionality of the bot

			case event.NewGetValMessageFromPeer:
				//log.Println("NewGetValMessageFromPeer received")

			default:
				log.Printf("Unhandled event: %v\n", message.EventType)
			}
		}

	}()

	// Block until a signal is received
	<-shutdown
	log.Println("Shutting down gracefully...")
}

func Unwrap(id int, onion string, msg string) *OverlayEnvelope {
	var envelope OverlayEnvelope
	err := json.Unmarshal([]byte(msg), &envelope)
	if err != nil {
		//		log.Errorf("json error: %v", err)
		return nil
	}

	envelope.ConversationID = id
	envelope.Onion = onion
	return &envelope
}

func getGroupHandle(name string) (string, int, error) {
	// Get the conversation list
	conversationList, err := utilities.Cwtchbot.Peer.FetchConversations()
	if err != nil {
		log.Printf("Failed to fetch conversations: %v", err)
		return "", 0, err
	}

	// Find the thread conversation
	for _, conversation := range conversationList {
		if conversation.Attributes != nil && conversation.Attributes["local.profile.name"] == name {
			log.Printf("Found matching conversation: handle: %s, ID: %d", conversation.Handle, conversation.ID)
			return conversation.Handle, conversation.ID, nil
		}
	}

	return "", 0, errors.New("no conversation found with matching thread name")
}

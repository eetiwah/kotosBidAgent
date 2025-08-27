package group

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"kotosBidAgent/agent/utilities"
	"log"
	"regexp"
	"strings"

	"cwtch.im/cwtch/model"
)

var debugPrintFlg = false

func InviteGroup(bundle string) string {
	// Is this a valid bundle
	name, err := decodeGroupInvite(bundle)
	if err != nil {
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, err.Error()))
	}

	err = utilities.Cwtchbot.Peer.ImportBundle(bundle)
	if err != nil {
		// Check if the error message indicates success
		if err.Error() != "importBundle.success" {
			return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, fmt.Sprintf("invite group process failed: %s, error: %v", bundle, err)))
		}

		// If the error message is "importBundle.success", treat it as a success
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, fmt.Sprintf("Invite Group was successful for: %s", name)))

	} else {
		return string(utilities.Cwtchbot.PackMessage(model.OverlayChat, "Error: this process path should have been used to add contact"))
	}

}

// decodeGroupInvite processes a cwtch group invite
func decodeGroupInvite(content string) (string, error) {
	// Log raw input
	log.Printf("\n\nRaw content: %q (length: %d)", content, len(content))

	// Split on ||
	parts := strings.Split(strings.TrimSpace(content), "||")
	log.Printf("Parts: %v (length: %d)", parts, len(parts))
	if len(parts) != 2 {
		return "", fmt.Errorf("error: invalid content: expected 2 parts, got %d", len(parts))
	}

	// Skip first 5 characters of parts[1]
	if len(parts[1]) <= 5 {
		return "", fmt.Errorf("error: parts[1] too short: %d chars", len(parts[1]))
	}
	encoded := parts[1][5:]
	log.Printf("Encoded Base64: %q (length: %d)", encoded, len(encoded))

	// Validate Base64
	if !regexp.MustCompile(`^[A-Za-z0-9+/=]+$`).MatchString(encoded) {
		return "", fmt.Errorf("error: invalid Base64 characters in encoded string")
	}

	// Decode Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		// Try RawStdEncoding for padding issues
		log.Printf("Error: StdEncoding failed: %v, trying RawStdEncoding", err)
		decodedBytes, err = base64.RawStdEncoding.DecodeString(encoded)
		if err != nil {
			return "", fmt.Errorf("error: base64 decode failed: %v", err)
		}
	}
	log.Printf("Decoded string: %s", string(decodedBytes))

	// Parse JSON
	var invite GroupInvite
	if err := json.Unmarshal(decodedBytes, &invite); err != nil {
		return "", fmt.Errorf("error: json decode failed: %v", err)
	}

	if debugPrintFlg {
		log.Printf("GroupID = %s\n", invite.GroupID)
		log.Printf("GroupName = %s\n", invite.GroupName)
		log.Printf("ServerHost = %s\n", invite.ServerHost)
		log.Printf("SharedKey = %s\n", invite.SharedKey)
	}

	return invite.GroupName, nil
}

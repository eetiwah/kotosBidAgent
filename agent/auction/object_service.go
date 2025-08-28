package auction

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kotosBidAgent/agent/utilities"
	"log"
	"net/http"
	"time"
)

func CreateAuctionObj(data []byte) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/createAuction", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		errMsg := fmt.Sprintf("Error: create_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: create_auction HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: create_auction HTTP status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	id, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Error: reading response body: %v", err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	log.Printf("AuctionId = %s", string(id))

	return nil
}

func StartAuction(data []byte) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/startAuction", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		errMsg := fmt.Sprintf("StopAuction: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("StopAuction: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("StopAuction: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func StopAuction(data []byte) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/stopAuction", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		errMsg := fmt.Sprintf("StopAuction: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("StopAuction: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("StopAuction: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func GetAuctionObj(id string) (AuctionObject, error) {
	var obj AuctionObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getAuction/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction HTTP Get %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_auction HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		errMsg := fmt.Sprintf("Error: get_auction decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return obj, err
	}

	return obj, nil
}

func GetAuctionList() ([]AuctionObject, error) {
	var _list []AuctionObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getAuctionList", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_auction_list HTTP status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("Error: get_auction_list decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return _list, err
	}

	return _list, nil
}

func GetBidObj(id string) (BidObject, error) {
	var obj BidObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getBid/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request with byteData as body
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_bid creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error: get_bid HTTP Get %s: %v", url, err)
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("Error: get_bid HTTP status code %d", resp.StatusCode)
		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return obj, errors.New(errMsg)
	}

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		errMsg := fmt.Sprintf("Error: get_bid decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return obj, err
	}

	return obj, nil
}

func GetBidList(id string) ([]BidObject, error) {
	var _list []BidObject

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/getBidList/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return _list, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetBidList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return _list, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetBidList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return _list, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&_list); err != nil {
		errMsg := fmt.Sprintf("GetStartAuctionList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		return _list, err
	}

	return _list, nil
}

// Not using the ones below

func Responsed(id string) (bool, error) {

	// Define the URL to the meta data manager
	url := fmt.Sprintf("%s/responded/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetAuctionList: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return false, err
	}
	//req.Header.Set("Content-Type", "application/json")

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetAuctionList: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http error", errMsg)
		return false, err
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetAuctionList: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return false, err
	}

	var _result string
	if err := json.NewDecoder(resp.Body).Decode(&_result); err != nil {
		errMsg := fmt.Sprintf("GetAuctionList: decoding assignments JSON: %v", err)
		log.Println(errMsg)
		//sendErrorMessage(conversationID, "Thread_Services: http status", errMsg)
		return false, err
	}

	// Have we responded already?
	if _result == "yes" {
		return true, nil
	} else {
		return false, nil
	}
}

func SetAuctionWinner(data []byte) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/setAuctionWinner", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionWinner: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("SetAuctionWinner: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("SetAuctionWinner: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func AddBid(obj BidObject) error {
	byteData, err := json.Marshal(obj)
	if err != nil {
		errMsg := fmt.Sprintf("AddBid: status code %d", err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/addBid", utilities.AUCTION_MGR_URI)

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewReader(byteData))
	if err != nil {
		errMsg := fmt.Sprintf("AddBid: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("AddBid: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("AddBid: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func OurBidId(id string) error {
	// Define the URL to the auction manager
	url := fmt.Sprintf("%s/isOurBid/%s", utilities.AUCTION_MGR_URI, id)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("GetOrder: creating HTTP request %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// Send request with timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("GetOrder: HTTP POST %s: %v", url, err)
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	defer resp.Body.Close()

	// Check HTTP status and read response body for errors
	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("GetOrder: status code %d", resp.StatusCode)

		if readErr == nil && len(body) > 0 {
			errMsg += fmt.Sprintf(": %s", string(body))
		}
		log.Println(errMsg)
		return errors.New(errMsg)
	}

	// We found the bidId, it is ours
	return nil
}

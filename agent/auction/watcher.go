package auction

func Watcher() {
	/*
		// Set tick interval
		tickInterval := time.Duration(10) * time.Second

		// Create a ticker that ticks every tickInterval
		ticker := time.NewTicker(tickInterval)

		// Define the function to be executed on each tick
		tickerFunc := func() error {

			// Get a list of active auctions that need to be responded to
			auctionList, err := GetAuctionList()
			if err != nil {
				return err
			}

			// Check to determine if there are any auctions that need to be responded to
			if len(auctionList) == 0 {
				log.Println("No auctions need to be responded to")
				return nil
			}

			// Respond to open auctions
			for _, auction := range auctionList {
				// Have we responded to the auction already
				respondFlg, err := Responsed(auction.AuctionId)
				if err != nil {
					log.Printf("Responded failed: %v", err)
					return err
				}

				if respondFlg {
					log.Printf("Already responded to auction = %s", auction.AuctionId)
					return nil
				}

				// Get Order from mongoDB
				order, err := GetOrder(auction.OrderId)
				if err != nil {
					log.Printf("GenerateBid failed: %v", err)
					return err
				}

				// Generate bid using order information
				bid, err := GenerateBid(auction.AuctionId, order)
				if err != nil {
					log.Printf("GenerateBid failed: %v", err)
					return err
				}

				// Save bid to mongoDB
				err = AddBid(bid)
				if err != nil {
					log.Printf("AddBid failed: %v", err)
					return err
				}

				// Send a bid to auction community
				err = Bid_Event_Sent(bid)
				if err != nil {
					log.Printf("Bid_Event_Sent failed: %v", err)
					return err
				}

				log.Printf("Auction: %s was responded to w/ bidId: %s\n", auction.AuctionId, bid.BidId)
			}

			return nil
		}

		// Create start message
		log.Println("Auction watcher started ...")

		// Create channels for completion and errors
		done := make(chan struct{})
		errChan := make(chan string, 1)

		// Start a goroutine to execute the function on each tick
		go func() {
			defer close(done)   // Ensure done is closed when the goroutine exits
			defer ticker.Stop() // Ensure ticker is stopped when the goroutine exits

			for range ticker.C {
				// Execute the function on each tick
				if err := tickerFunc(); err != nil {
					errChan <- "Watcher Error:  " + err.Error()
					return
				}
			}
		}()

		// Wait for the goroutine to complete or fail
		select {
		case <-done:
			// Check if an error occurred
			select {
			case errMsg := <-errChan:
				log.Printf("Error = %v", errMsg)

			default:
				log.Println("Inventory check completed")
			}
		case errMsg := <-errChan:
			log.Printf("Error = %v", errMsg)
		}
	*/
}

/*
func GenerateBid(auctionId string, obj OrderObject) (BidObject, error) {
	bid := BidObject{
		BidId:        uuid.NewString(),
		AuctionId:    auctionId,
		Price:        obj.Price,
		Quantity:     obj.Quantity,
		DeliveryDate: obj.DeliveryDate,
		Onion:        utilities.EntityOnion,
		ResponseDate: time.Now().UTC(),
	}

	return bid, nil
}
*/

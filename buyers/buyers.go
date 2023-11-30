package buyers

import (
	items "groww/Items"
	"groww/sellers"

	"github.com/google/uuid"
)

func init() {
	buyersMap = make(map[string]*Buyer)
}

var buyersMap map[string]*Buyer

type Buyer struct {
	ID        string
	Purchases []*items.Item
	Wishlist  []*items.Item
}

func AddBuyer() *Buyer {
	buyerID := uuid.NewString()
	buyer := &Buyer{
		ID: buyerID,
	}
	buyersMap[buyerID] = buyer
	return buyer
}

func (buyer *Buyer) SeeSellerListings(sellerId string) []*items.Item {
	seller := sellers.FindSeller(sellerId)
	if seller == nil {
		return nil
	}
	return seller.Items
}

func (buyer *Buyer) Buy(itemID string) *items.Item {
	item := items.GetItemByID(itemID)
	if item == nil || item.Sold {
		return nil
	}
	item.Buy(buyer.ID)
	buyer.Purchases = append(buyer.Purchases, item)
	return item
}

func (buyer *Buyer) AddToWishlist(itemID string) *items.Item {
	item := items.GetItemByID(itemID)
	if item.Sold {
		return nil
	}
	buyer.Wishlist = append(buyer.Wishlist, item)
	return item
}

func FetchBuyerByID(buyerID string) *Buyer {
	buyer, ok := buyersMap[buyerID]
	if ok {
		return buyer
	}
	return nil
}

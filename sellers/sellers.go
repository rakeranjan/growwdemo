package sellers

import (
	items "groww/Items"

	"github.com/google/uuid"
)

func init() {
	sellerMap = make(map[string]*Seller)
}

var sellerMap map[string]*Seller

type Seller struct {
	ID    string
	Items []*items.Item
}

func AddSeller() *Seller {
	id := uuid.NewString()
	seller := &Seller{
		ID:    id,
		Items: make([]*items.Item, 0),
	}
	sellerMap[id] = seller
	return seller
}

func (seller *Seller) AddItem(name string, price int) *items.Item {
	item := items.AddItem(name, seller.ID, price)
	seller.Items = append(seller.Items, item)
	return item
}

func (seller *Seller) GetAllListings() []*items.Item {
	return seller.Items
}

func FindSeller(sellerID string) *Seller {
	seller, ok := sellerMap[sellerID]
	if ok {
		return seller
	}
	return nil
}

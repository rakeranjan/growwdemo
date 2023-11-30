package items

import "github.com/google/uuid"

func init() {
	itemMap = make(map[string]*Item)
}

var itemMap map[string]*Item

type Item struct {
	ID     string
	Name   string
	Price  int
	Seller string
	buyer  string
	Sold   bool
}

func AddItem(name, seller string, price int) *Item {
	id := uuid.NewString()
	item := &Item{
		ID:     id,
		Name:   name,
		Seller: seller,
		Price:  price,
		Sold:   false,
	}
	itemMap[id] = item
	return item
}

func (item *Item) Buy(buyer string) {
	item.Sold = true
	item.buyer = buyer
}

func GetItemByID(id string) *Item {
	item, ok := itemMap[id]
	if ok {
		return item
	}
	return nil
}

func GetAllItem(sold bool) []*Item {
	items := make([]*Item, 0)
	for _, v := range itemMap {
		if v.Sold == sold {
			items = append(items, v)
		}
	}
	if len(items) == 0 {
		return nil
	}
	return items
}

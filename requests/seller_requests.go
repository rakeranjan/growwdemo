package requests

type ListItem struct {
	SellerID string `json:"sellerId"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}

type GetSellerListingsRequests struct {
	SellerID string `json:"sellerId"`
}

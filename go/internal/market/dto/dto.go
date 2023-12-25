package dto

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetID       string  `json:"asset_id"`
	CurrentShares int     `json:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderID           string               `json:"order_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         string               `json:"order_type"`
	Status            string               `json:"status"`
	Partial           string               `json:"partial"`
	Shares            string               `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transation"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transacrion_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}

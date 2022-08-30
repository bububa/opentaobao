package car

// SellerInfo 结构体
type SellerInfo struct {
	// sellerEmail
	SellerEmail string `json:"seller_email,omitempty" xml:"seller_email,omitempty"`
	// sellerNick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// sellerPhone
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// sellerId
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

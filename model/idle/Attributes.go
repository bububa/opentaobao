package idle

// Attributes 结构体
type Attributes struct {
	// 用户寄出快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 卖家收件地址
	SellerAddressDetail string `json:"seller_address_detail,omitempty" xml:"seller_address_detail,omitempty"`
	// 卖家收件电话
	SellerAddressPhone string `json:"seller_address_phone,omitempty" xml:"seller_address_phone,omitempty"`
	// 卖家人收件名
	SellerAddressName string `json:"seller_address_name,omitempty" xml:"seller_address_name,omitempty"`
}

package traveltrade

// BuyerInfo 结构体
type BuyerInfo struct {
	// 买家邮件地址
	BuyerEmail string `json:"buyer_email,omitempty" xml:"buyer_email,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 买家联系方式
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 买家是否已评价
	BuyerRate bool `json:"buyer_rate,omitempty" xml:"buyer_rate,omitempty"`
}

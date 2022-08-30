package user

// OpenUidInfo 结构体
type OpenUidInfo struct {
	// 买家openuid
	BuyerOpenUid string `json:"buyer_open_uid,omitempty" xml:"buyer_open_uid,omitempty"`
	// 买家nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

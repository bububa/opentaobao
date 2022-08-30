package promotion

// ErrorMessage 结构体
type ErrorMessage struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 发送失败的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// asas
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

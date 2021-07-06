package servicecenter

// SubUser 结构体
type SubUser struct {
	// 授权状态描述
	StateDes string `json:"state_des,omitempty" xml:"state_des,omitempty"`
	// 商家子账号昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 授权状态: 1表示授权，0表示取消授权，-1表示失效
	State int64 `json:"state,omitempty" xml:"state,omitempty"`
}

package cainiaohandover

// HandoverContentUpdateParcelDto 结构体
type HandoverContentUpdateParcelDto struct {
	// 小包对应的店铺账号;比如cnxxxx;填入补充相关信息性能更好
	LoginId string `json:"login_id,omitempty" xml:"login_id,omitempty"`
	// 小包的LP号,必填;
	LpCode string `json:"lp_code,omitempty" xml:"lp_code,omitempty"`
	// 小包对应的店铺id;填入相关信息性能更好
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

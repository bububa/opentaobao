package nrt

// StallSigningReqDto 结构体
type StallSigningReqDto struct {
	// 银行卡信息
	BizCards []SettleCardInfoDto `json:"biz_cards,omitempty" xml:"biz_cards>settle_card_info_dto,omitempty"`
	// 对应卖场的smid
	IpRoleId string `json:"ip_role_id,omitempty" xml:"ip_role_id,omitempty"`
	// 摊位名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

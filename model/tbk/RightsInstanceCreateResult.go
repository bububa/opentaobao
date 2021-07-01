package tbk

// RightsInstanceCreateResult 结构体
type RightsInstanceCreateResult struct {
	// 淘礼金Id
	RightsId string `json:"rights_id,omitempty" xml:"rights_id,omitempty"`
	// 淘礼金领取Url
	SendUrl string `json:"send_url,omitempty" xml:"send_url,omitempty"`
	// 投放code【百川商品详情页业务专用】
	VegasCode string `json:"vegas_code,omitempty" xml:"vegas_code,omitempty"`
	// 创建完成后资金账户可用资金，单位元，保留2位小数
	AvailableFee string `json:"available_fee,omitempty" xml:"available_fee,omitempty"`
}

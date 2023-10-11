package simba

// TopMarketSceneVo 结构体
type TopMarketSceneVo struct {
	// 场景code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 场景名称
	MarketName string `json:"market_name,omitempty" xml:"market_name,omitempty"`
	// code 对应的 appid
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

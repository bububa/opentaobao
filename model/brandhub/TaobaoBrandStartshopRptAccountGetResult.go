package brandhub

// TaobaoBrandStartshopRptAccountGetResult 结构体
type TaobaoBrandStartshopRptAccountGetResult struct {
	// thedate
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 展现量
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击量
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 花费(元)
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 千次展现成本
	Cpm string `json:"cpm,omitempty" xml:"cpm,omitempty"`
	// 点击单价
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 宝贝收藏数
	Favitemtotal int64 `json:"favitemtotal,omitempty" xml:"favitemtotal,omitempty"`
	// 展示成交笔数
	Transactionshippingtotal int64 `json:"transactionshippingtotal,omitempty" xml:"transactionshippingtotal,omitempty"`
	// 店铺收藏数
	Favshoptotal int64 `json:"favshoptotal,omitempty" xml:"favshoptotal,omitempty"`
	// 宝贝加购数
	Carttotal int64 `json:"carttotal,omitempty" xml:"carttotal,omitempty"`
	// 展示成交金额
	Transactiontotal string `json:"transactiontotal,omitempty" xml:"transactiontotal,omitempty"`
	// 展示回报率
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 展示转化率
	Cvr string `json:"cvr,omitempty" xml:"cvr,omitempty"`
	// 点击成交笔数
	ClickTransactionshipping int64 `json:"click_transactionshipping,omitempty" xml:"click_transactionshipping,omitempty"`
	// 点击成交金额
	ClickTransactiontotal string `json:"click_transactiontotal,omitempty" xml:"click_transactiontotal,omitempty"`
	// 点击回报率
	ClickRoi string `json:"click_roi,omitempty" xml:"click_roi,omitempty"`
	// 点击转化率
	ClickCvr string `json:"click_cvr,omitempty" xml:"click_cvr,omitempty"`
	// 点击访客数
	ClickUv int64 `json:"click_uv,omitempty" xml:"click_uv,omitempty"`
	// 触达访客数
	Uv string `json:"uv,omitempty" xml:"uv,omitempty"`
	// 触达新访客数
	UvNew string `json:"uv_new,omitempty" xml:"uv_new,omitempty"`
}

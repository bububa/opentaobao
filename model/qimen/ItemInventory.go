package qimen

// ItemInventory 结构体
type ItemInventory struct {
	// 奇门仓储字段,C123,string(50),,
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	ChannelCode string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 奇门仓储字段,C123,string(50),,
	LockQuantity string `json:"lockQuantity,omitempty" xml:"lockQuantity,omitempty"`
	// 奇门仓储字段
	OrderSourceCode string `json:"orderSourceCode,omitempty" xml:"orderSourceCode,omitempty"`
	// 奇门仓储字段
	SubSourceCode string `json:"subSourceCode,omitempty" xml:"subSourceCode,omitempty"`
	// 奇门仓储字段
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	CombItemId string `json:"combItemId,omitempty" xml:"combItemId,omitempty"`
	// 奇门仓储字段
	ItemQuantity string `json:"itemQuantity,omitempty" xml:"itemQuantity,omitempty"`
}

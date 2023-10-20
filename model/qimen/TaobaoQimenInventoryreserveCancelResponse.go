package qimen

// TaobaoqimeninventoryreservecancelResponse 结构体
type TaobaoqimeninventoryreservecancelResponse struct {
	// 奇门仓储字段
	ItemInventories []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 奇门仓储字段
	IsRetry string `json:"isRetry,omitempty" xml:"isRetry,omitempty"`
}

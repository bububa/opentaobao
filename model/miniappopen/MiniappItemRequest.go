package miniappopen

// MiniappItemRequest 结构体
type MiniappItemRequest struct {
	// 商品id列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 小程序小游戏appId
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

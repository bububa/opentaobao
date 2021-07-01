package qimen

// TaobaoQimenShopSynchronizeRequest 结构体
type TaobaoQimenShopSynchronizeRequest struct {
	// add,update, 必填
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 店铺
	Shop *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
}

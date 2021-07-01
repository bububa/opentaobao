package qimen

// BatchItemSynItem 结构体
type BatchItemSynItem struct {
	// 没有同步成功的商品的编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 出错信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

package logistic

// LogisticsResourceRequest 结构体
type LogisticsResourceRequest struct {
	// 可选值:offline(自己联系发货),online(在线下单),all(自己联系+在线下单)instant(同城配送).
	OrderMode string `json:"order_mode,omitempty" xml:"order_mode,omitempty"`
}

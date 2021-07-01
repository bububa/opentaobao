package ieagency

// VirProOrderVo 结构体
type VirProOrderVo struct {
	// bookTime
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

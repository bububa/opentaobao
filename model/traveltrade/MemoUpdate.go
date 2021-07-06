package traveltrade

// MemoUpdate 结构体
type MemoUpdate struct {
	// 关闭订单时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 交易ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

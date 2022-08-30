package tmallgenie

// PurchaseCircleInfoForOuterDto 结构体
type PurchaseCircleInfoForOuterDto struct {
	// 圈选内容
	CircleContent string `json:"circle_content,omitempty" xml:"circle_content,omitempty"`
	// 圈选标识id,由采销系统分配
	CircleId int64 `json:"circle_id,omitempty" xml:"circle_id,omitempty"`
}

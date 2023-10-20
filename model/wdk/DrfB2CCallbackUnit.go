package wdk

// DrfB2ccallbackUnit 结构体
type DrfB2ccallbackUnit struct {
	// 作业内容
	CallbackContents []DrfB2ccallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_b2ccallback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

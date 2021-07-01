package wdk

// DrfB2CCallbackUnit 结构体
type DrfB2CCallbackUnit struct {
	// 作业内容
	CallbackContents []DrfB2CCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_b2c_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

package wdk

// SfB2cfmsCallbackUnit 结构体
type SfB2cfmsCallbackUnit struct {
	// 作业内容
	CallbackContents []SfB2cfmsCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>sf_b2cfms_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

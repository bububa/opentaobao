package wdk

// MissfreshO2OCallbackUnit 结构体
type MissfreshO2OCallbackUnit struct {
	// 作业内容
	CallbackContents []MissfreshO2OCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>missfresh_o2o_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

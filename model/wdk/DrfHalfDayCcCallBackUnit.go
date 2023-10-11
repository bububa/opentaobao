package wdk

// DrfHalfDayCcCallbackUnit 结构体
type DrfHalfDayCcCallbackUnit struct {
	// 作业内容
	CallbackContents []DrfHalfDayCcCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_half_day_cc_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
	// 作业单元扩展属性
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

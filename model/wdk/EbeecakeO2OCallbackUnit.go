package wdk

// EbeecakeO2ocallbackUnit 结构体
type EbeecakeO2ocallbackUnit struct {
	// 作业内容列表
	CallbackContents []EbeecakeO2ocallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>ebeecake_o2ocallback_content,omitempty"`
	// 作业单元号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

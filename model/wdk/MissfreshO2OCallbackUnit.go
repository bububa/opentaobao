package wdk

// MissfreshO2ocallbackUnit 结构体
type MissfreshO2ocallbackUnit struct {
	// 作业内容
	CallbackContents []MissfreshO2ocallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>missfresh_o2ocallback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

package wdk

// DrfHalfDayCcCallBackUnit 
type DrfHalfDayCcCallBackUnit struct {
    // 作业内容
    CallbackContents   []DrfHalfDayCcCallBackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_half_day_cc_call_back_content,omitempty"`
    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
    // 作业单元扩展属性
    Attribute   string `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

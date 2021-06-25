package wdk

// MissfreshO2OCallbackUnit 
type MissfreshO2OCallbackUnit struct {

    // 作业内容
    CallbackContents   []MissfreshO2OCallbackContent `json:"callback_contents,omitempty"`

    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

}

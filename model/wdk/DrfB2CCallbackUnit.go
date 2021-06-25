package wdk

// DrfB2CCallbackUnit 
type DrfB2CCallbackUnit struct {

    // 作业内容
    CallbackContents   []DrfB2CCallbackContent `json:"callback_contents,omitempty"`

    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

}

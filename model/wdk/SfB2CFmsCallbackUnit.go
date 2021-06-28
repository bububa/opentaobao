package wdk

// SfB2CFmsCallbackUnit 
type SfB2CFmsCallbackUnit struct {

    // 作业内容
    
    CallbackContents   []SfB2CFmsCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents,omitempty"`
    

    // 作业单元单号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
    

}

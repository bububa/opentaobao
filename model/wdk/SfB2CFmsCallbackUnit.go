package wdk

// SfB2CFmsCallbackUnit 
/* model for simplify = false
type SfB2CFmsCallbackUnit struct {

    // 作业内容
    
    CallbackContents  struct {
        SfB2CFmsCallbackContent  []SfB2CFmsCallbackContent `json:"sf_b2c_fms_callback_content,omitempty"`
    } `json:"callback_contents,omitempty"`
    

    // 作业单元单号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`
    

}
*/

// SfB2CFmsCallbackUnit 
type SfB2CFmsCallbackUnit struct {

    // 作业内容
    CallbackContents   []SfB2CFmsCallbackContent `json:"callback_contents,omitempty"`

    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

}

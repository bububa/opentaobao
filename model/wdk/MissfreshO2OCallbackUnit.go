package wdk

// MissfreshO2OCallbackUnit 
/* model for simplify = false
type MissfreshO2OCallbackUnit struct {

    // 作业内容
    
    CallbackContents  struct {
        MissfreshO2OCallbackContent  []MissfreshO2OCallbackContent `json:"missfresh_o2o_callback_content,omitempty"`
    } `json:"callback_contents,omitempty"`
    

    // 作业单元单号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`
    

}
*/

// MissfreshO2OCallbackUnit 
type MissfreshO2OCallbackUnit struct {

    // 作业内容
    CallbackContents   []MissfreshO2OCallbackContent `json:"callback_contents,omitempty"`

    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

}

package wdk

// EbeecakeO2OCallbackUnit 
/* model for simplify = false
type EbeecakeO2OCallbackUnit struct {

    // 作业内容列表
    
    CallbackContents  struct {
        EbeecakeO2OCallbackContent  []EbeecakeO2OCallbackContent `json:"ebeecake_o2o_callback_content,omitempty"`
    } `json:"callback_contents,omitempty"`
    

    // 作业单元号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`
    

}
*/

// EbeecakeO2OCallbackUnit 
type EbeecakeO2OCallbackUnit struct {

    // 作业内容列表
    CallbackContents   []EbeecakeO2OCallbackContent `json:"callback_contents,omitempty"`

    // 作业单元号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

}

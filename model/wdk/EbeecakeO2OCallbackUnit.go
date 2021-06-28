package wdk

// EbeecakeO2OCallbackUnit 
type EbeecakeO2OCallbackUnit struct {

    // 作业内容列表
    
    CallbackContents   []EbeecakeO2OCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents,omitempty"`
    

    // 作业单元号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
    

}

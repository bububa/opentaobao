package wdk

// DrfHalfDayCcCallBackUnit 
/* model for simplify = false
type DrfHalfDayCcCallBackUnit struct {

    // 作业内容
    
    CallbackContents  struct {
        DrfHalfDayCcCallBackContent  []DrfHalfDayCcCallBackContent `json:"drf_half_day_cc_call_back_content,omitempty"`
    } `json:"callback_contents,omitempty"`
    

    // 作业单元单号
    
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`
    

    // 作业单元扩展属性
    
    Attribute   string `json:"attribute,omitempty"`
    

}
*/

// DrfHalfDayCcCallBackUnit 
type DrfHalfDayCcCallBackUnit struct {

    // 作业内容
    CallbackContents   []DrfHalfDayCcCallBackContent `json:"callback_contents,omitempty"`

    // 作业单元单号
    WorkOrderUnitId   string `json:"work_order_unit_id,omitempty"`

    // 作业单元扩展属性
    Attribute   string `json:"attribute,omitempty"`

}

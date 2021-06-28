package wdk

// SfB2CFmsCallbackOrder 
type SfB2CFmsCallbackOrder struct {

    // 作业状态变更时间
    
    StatusChangeTime   string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
    

    // 作业状态变更类型： START_PICK(“开始拣货”)， PICK_FINISH(“拣货完成”)， START_PACKAGE(“开始打包”), PACKAGE _FINISH(“打包完成”);
    
    StatusChangeType   string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
    

    // 节点编码
    
    NodeCode   string `json:"node_code,omitempty" xml:"node_code,omitempty"`
    

    // 作业单号
    
    WorkOrderId   string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
    

    // 作业单元
    
    CallbackUnits   []SfB2CFmsCallbackUnit `json:"callback_units,omitempty" xml:"callback_units,omitempty"`
    

    // 操作员
    
    Operator   *Operator `json:"operator,omitempty" xml:"operator,omitempty"`
    

}

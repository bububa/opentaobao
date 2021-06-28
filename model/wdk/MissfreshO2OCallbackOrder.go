package wdk

// MissfreshO2OCallbackOrder 
/* model for simplify = false
type MissfreshO2OCallbackOrder struct {

    // 作业状态变更时间
    
    StatusChangeTime   string `json:"status_change_time,omitempty"`
    

    // 作业状态变更类型： START_PICK(“开始拣货”)， PICK_FINISH(“拣货完成”)， START_PACKAGE(“开始打包”), PACKAGE _FINISH(“打包完成”);
    
    StatusChangeType   string `json:"status_change_type,omitempty"`
    

    // 节点编码
    
    NodeCode   string `json:"node_code,omitempty"`
    

    // 作业单类型： BATCH("批次"),  ORDER("物流单
    
    WorkOrderType   string `json:"work_order_type,omitempty"`
    

    // 作业单号
    
    WorkOrderId   string `json:"work_order_id,omitempty"`
    

    // 作业单元
    
    CallbackUnits  struct {
        MissfreshO2OCallbackUnit  []MissfreshO2OCallbackUnit `json:"missfresh_o2o_callback_unit,omitempty"`
    } `json:"callback_units,omitempty"`
    

    // 是否作业节点终态
    
    IsFinal   bool `json:"is_final,omitempty"`
    

    // 操作员
    
    Operator  *struct {
        Operator  *Operator `json:"operator,omitempty"`
    } `json:"operator,omitempty"`
    

    // 容器列表
    
    Containers  struct {
        Container  []Container `json:"container,omitempty"`
    } `json:"containers,omitempty"`
    

}
*/

// MissfreshO2OCallbackOrder 
type MissfreshO2OCallbackOrder struct {

    // 作业状态变更时间
    StatusChangeTime   string `json:"status_change_time,omitempty"`

    // 作业状态变更类型： START_PICK(“开始拣货”)， PICK_FINISH(“拣货完成”)， START_PACKAGE(“开始打包”), PACKAGE _FINISH(“打包完成”);
    StatusChangeType   string `json:"status_change_type,omitempty"`

    // 节点编码
    NodeCode   string `json:"node_code,omitempty"`

    // 作业单类型： BATCH("批次"),  ORDER("物流单
    WorkOrderType   string `json:"work_order_type,omitempty"`

    // 作业单号
    WorkOrderId   string `json:"work_order_id,omitempty"`

    // 作业单元
    CallbackUnits   []MissfreshO2OCallbackUnit `json:"callback_units,omitempty"`

    // 是否作业节点终态
    IsFinal   bool `json:"is_final,omitempty"`

    // 操作员
    Operator   *Operator `json:"operator,omitempty"`

    // 容器列表
    Containers   []Container `json:"containers,omitempty"`

}

package wdk

// EbeecakeO2OCallbackOrder 
type EbeecakeO2OCallbackOrder struct {
    // 作业单号
    WorkOrderId   string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
    // 作业单类型： BATCH("批次"), ORDER("物流单")
    WorkOrderType   string `json:"work_order_type,omitempty" xml:"work_order_type,omitempty"`
    // 作业状态变更类型：SHIP("揽收"),SIGN("妥投"),SIGN_ERROR("配送异常"),REFUSE("拒收")
    StatusChangeType   string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
    // 作业状态变更时间
    StatusChangeTime   string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
    // 作业单元列表
    CallbackUnits   []EbeecakeO2OCallbackUnit `json:"callback_units,omitempty" xml:"callback_units>ebeecake_o2o_callback_unit,omitempty"`
    // 配送员
    Postman   *Postman `json:"postman,omitempty" xml:"postman,omitempty"`
}

package wdk

// DeliveryCallbackOrder 
/* model for simplify = false
type DeliveryCallbackOrder struct {

    // 作业单号
    
    WorkOrderId   string `json:"work_order_id,omitempty"`
    

    // 作业状态变更类型：SHIP("揽收"),SIGN("妥投"),SIGN_ERROR("配送异常"),REFUSE("拒收")
    
    StatusChangeType   string `json:"status_change_type,omitempty"`
    

    // 作业状态变更时间
    
    StatusChangeTime   string `json:"status_change_time,omitempty"`
    

    // 拒收子单列表
    
    RefusedOrderDetails  struct {
        DeliveryCallbackOrderDetail  []DeliveryCallbackOrderDetail `json:"delivery_callback_order_detail,omitempty"`
    } `json:"refused_order_details,omitempty"`
    

    // 配送员
    
    Deliveryman  *struct {
        Deliveryman  *Deliveryman `json:"deliveryman,omitempty"`
    } `json:"deliveryman,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 配送站编码
    
    DeliveryDockCode   string `json:"delivery_dock_code,omitempty"`
    

    // 来源系统:：CHINA_POST:邮政
    
    SourceSystem   string `json:"source_system,omitempty"`
    

}
*/

// DeliveryCallbackOrder 
type DeliveryCallbackOrder struct {

    // 作业单号
    WorkOrderId   string `json:"work_order_id,omitempty"`

    // 作业状态变更类型：SHIP("揽收"),SIGN("妥投"),SIGN_ERROR("配送异常"),REFUSE("拒收")
    StatusChangeType   string `json:"status_change_type,omitempty"`

    // 作业状态变更时间
    StatusChangeTime   string `json:"status_change_time,omitempty"`

    // 拒收子单列表
    RefusedOrderDetails   []DeliveryCallbackOrderDetail `json:"refused_order_details,omitempty"`

    // 配送员
    Deliveryman   *Deliveryman `json:"deliveryman,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 配送站编码
    DeliveryDockCode   string `json:"delivery_dock_code,omitempty"`

    // 来源系统:：CHINA_POST:邮政
    SourceSystem   string `json:"source_system,omitempty"`

}

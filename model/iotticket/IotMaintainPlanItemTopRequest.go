package iotticket

// IotMaintainPlanItemTopRequest 
type IotMaintainPlanItemTopRequest struct {

    // 付款角色：merchant-商家记账；customer-客户支付
    
    PayRole   string `json:"pay_role,omitempty" xml:"pay_role,omitempty"`
    

    // 设备编码（需要映射）
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

}

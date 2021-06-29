package middleclaims

// ClaimsAcceptDto 
type ClaimsAcceptDto struct {

    // 服务工单ID
    
    ServiceWorkOrderId   int64 `json:"service_work_order_id,omitempty" xml:"service_work_order_id,omitempty"`
    

    // 主订单ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 子订单ID
    
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

    // 理赔受理结果
    
    ClaimAcceptResult   bool `json:"claim_accept_result,omitempty" xml:"claim_accept_result,omitempty"`
    

    // 理赔受理拒绝原因
    
    ClaimsAcceptResultDesc   string `json:"claims_accept_result_desc,omitempty" xml:"claims_accept_result_desc,omitempty"`
    

    // 报案号
    
    ReportNo   string `json:"report_no,omitempty" xml:"report_no,omitempty"`
    

    // 保单号
    
    PolicyNo   string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
    

    // 预留扩展Map
    
    ExtensionMap   *Extensionmap `json:"extension_map,omitempty" xml:"extension_map,omitempty"`
    

    // 保险退货仓库地址实体类
    
    DeliveryAddressDto   *DeliveryAddressDto `json:"delivery_address_dto,omitempty" xml:"delivery_address_dto,omitempty"`
    

}

package medicalbase

// OrderlSyncDto 
type OrderlSyncDto struct {

    // 登录用户支付宝ID
    
    AlipayId   string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
    

    // 医院唯一标识
    
    HosOrgNo   string `json:"hos_org_no,omitempty" xml:"hos_org_no,omitempty"`
    

    // 院区编码
    
    HosDistinctCode   string `json:"hos_distinct_code,omitempty" xml:"hos_distinct_code,omitempty"`
    

    // 预约ID
    
    ReservationId   string `json:"reservation_id,omitempty" xml:"reservation_id,omitempty"`
    

    // 订单状态
    
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 失败原因
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 取号密码
    
    OrderPass   string `json:"order_pass,omitempty" xml:"order_pass,omitempty"`
    

}

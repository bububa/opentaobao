package normalvisa

// NormalVisaUpdateUnit 
type NormalVisaUpdateUnit struct {

    // 预约时间
    
    BookTime   string `json:"book_time,omitempty" xml:"book_time,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 状态：无办理人信息（1001），办理人已填写（1002），已收到资料（1003），已审核完成（1004），已送签（1005），结果已返回（1006），已预约面试（1007  ），处理中（1008），买家已填写反馈信息（1009），已中止办理（1010）
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 物流单号
    
    PostNumber   string `json:"post_number,omitempty" xml:"post_number,omitempty"`
    

    // 订单号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 用户id
    
    PersonVisaId   int64 `json:"person_visa_id,omitempty" xml:"person_visa_id,omitempty"`
    

    // 预约地点
    
    BookPlace   string `json:"book_place,omitempty" xml:"book_place,omitempty"`
    

    // 物流公司编码
    
    PostCompanyCode   string `json:"post_company_code,omitempty" xml:"post_company_code,omitempty"`
    

    // 物流公司名称
    
    PostCompanyName   string `json:"post_company_name,omitempty" xml:"post_company_name,omitempty"`
    

}

package lstlogistics

// SendOfflineOrderParam 
type SendOfflineOrderParam struct {

    // 快递单号
    
    MailNo   string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
    

    // 物流公司code
    
    CpCompanyCode   string `json:"cp_company_code,omitempty" xml:"cp_company_code,omitempty"`
    

    // 物流公司名称
    
    CpCompanyName   string `json:"cp_company_name,omitempty" xml:"cp_company_name,omitempty"`
    

    // 发货时间
    
    SendTime   string `json:"send_time,omitempty" xml:"send_time,omitempty"`
    

    // 备注
    
    Remarks   string `json:"remarks,omitempty" xml:"remarks,omitempty"`
    

    // 发货主订单列表
    
    MainOrderParamList   []MainOrderParam `json:"main_order_param_list,omitempty" xml:"main_order_param_list,omitempty"`
    

}

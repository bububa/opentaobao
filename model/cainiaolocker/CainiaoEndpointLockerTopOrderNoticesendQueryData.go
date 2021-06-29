package cainiaolocker

// CainiaoEndpointLockerTopOrderNoticesendQueryData 
type CainiaoEndpointLockerTopOrderNoticesendQueryData struct {

    // 快递公司编码
    
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    

    // 用于返回淘系包裹脱密手机号用作数据核对
    
    GetterPhone   string `json:"getter_phone,omitempty" xml:"getter_phone,omitempty"`
    

    // 是否需要输入手机号，false-不需要，裹裹可以自己判断手机号，true-需要手动输入手机号
    
    NeedInputPhone   bool `json:"need_input_phone,omitempty" xml:"need_input_phone,omitempty"`
    

    // 裹裹发送通知消息标识，false-?非裹裹发送，true-裹裹发送
    
    GuoguoSendNoticeFlag   bool `json:"guoguo_send_notice_flag,omitempty" xml:"guoguo_send_notice_flag,omitempty"`
    

    // 快递公司名称
    
    CpName   string `json:"cp_name,omitempty" xml:"cp_name,omitempty"`
    

}

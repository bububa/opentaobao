package examination

// CooperationOrderInfo 
type CooperationOrderInfo struct {

    // 机构checkNo
    
    CheckNo   string `json:"check_no,omitempty" xml:"check_no,omitempty"`
    

    // 证件号
    
    CertNumber   string `json:"cert_number,omitempty" xml:"cert_number,omitempty"`
    

    // 电话号码
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

    // 婚否(0-未婚; 1-已婚)
    
    Married   string `json:"married,omitempty" xml:"married,omitempty"`
    

    // 证件类型(0-身份证; 1-护照; 2-军官证)
    
    CertType   string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
    

    // 体检人姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 性别(0-男;1-女;)
    
    Gender   string `json:"gender,omitempty" xml:"gender,omitempty"`
    

    // 体检日期
    
    CheckDate   string `json:"check_date,omitempty" xml:"check_date,omitempty"`
    

    // 体检状态：未到检(exam_not), 已到检(exam_done)，已取消（exam_cancel）(新增)； 上门服务中还需以下两种状态：预约确认中（reserve_confirming），预约拒绝（reserve_rejected）；
    
    Reportstatus   string `json:"reportstatus,omitempty" xml:"reportstatus,omitempty"`
    

    // 机构的预约唯一码
    
    UniqReserveCode   string `json:"uniq_reserve_code,omitempty" xml:"uniq_reserve_code,omitempty"`
    

    // 阿里健康预约唯一标识
    
    ReserveNumber   string `json:"reserve_number,omitempty" xml:"reserve_number,omitempty"`
    

    // 预约拒绝原因，在预约拒绝后才需赋值（仅上门服务中使用）
    
    CancelReason   string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
    

}

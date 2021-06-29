package alihealthmedical

// ServiceResult 
type ServiceResult struct {

    // 返回数据对象
    
    Data   *OuterMsgPullVo `json:"data,omitempty" xml:"data,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errMessage
    
    ErrMessage   string `json:"err_message,omitempty" xml:"err_message,omitempty"`
    

    // 审核记录通讯id
    
    DoctorUuid   string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
    

    // 消息是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    

    // 错误码
    
    CodeError   string `json:"code_error,omitempty" xml:"code_error,omitempty"`
    

    // 错误信息
    
    MessageError   string `json:"message_error,omitempty" xml:"message_error,omitempty"`
    

}

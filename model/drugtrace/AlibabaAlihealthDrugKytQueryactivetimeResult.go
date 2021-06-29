package drugtrace

// AlibabaAlihealthDrugKytQueryactivetimeResult 
type AlibabaAlihealthDrugKytQueryactivetimeResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 码激活状态DTO
    
    Models   []CodeActiveStatusDto `json:"models,omitempty" xml:"models,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

}

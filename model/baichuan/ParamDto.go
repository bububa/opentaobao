package baichuan

// ParamDto 
type ParamDto struct {

    // 业务参数，传递需要判断的口令
    
    BizParam   string `json:"biz_param,omitempty" xml:"biz_param,omitempty"`
    

    // 扩展参数
    
    ExtraParam   string `json:"extra_param,omitempty" xml:"extra_param,omitempty"`
    

    // 系统自动生成
    
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    

}

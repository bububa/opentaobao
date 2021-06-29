package campus

// WorkBenchContext 
type WorkBenchContext struct {

    // 应用id
    
    SystemId   string `json:"system_id,omitempty" xml:"system_id,omitempty"`
    

    // 公司id
    
    CompanyId   int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
    

    // 园区id
    
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    

    // app编码
    
    AppCode   string `json:"app_code,omitempty" xml:"app_code,omitempty"`
    

    // securityCode
    
    SecurityCode   string `json:"security_code,omitempty" xml:"security_code,omitempty"`
    

    // userId
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // campusCode
    
    CampusCode   string `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
    

    // userName
    
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    

    // eagleEyeTraceId
    
    EagleEyeTraceId   string `json:"eagle_eye_trace_id,omitempty" xml:"eagle_eye_trace_id,omitempty"`
    

    // ip
    
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
    

    // 登录者的语言
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

}

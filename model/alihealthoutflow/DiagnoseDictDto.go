package alihealthoutflow

// DiagnoseDictDto 
type DiagnoseDictDto struct {

    // icd10编码(非空)
    
    IcdCode   string `json:"icd_code,omitempty" xml:"icd_code,omitempty"`
    

    // icd10名称(非空)
    
    IcdName   string `json:"icd_name,omitempty" xml:"icd_name,omitempty"`
    

    // his诊断编码(非空)
    
    HisDiagnoseCode   string `json:"his_diagnose_code,omitempty" xml:"his_diagnose_code,omitempty"`
    

    // his诊断名称(非空)
    
    HisDiagnoseName   string `json:"his_diagnose_name,omitempty" xml:"his_diagnose_name,omitempty"`
    

    // 1可用0停用(非空)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 渠道(非空)
    
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
    

}

package alihealthoutflow

// UsageRequest 
type UsageRequest struct {

    // 用法code(非空)
    
    HisUsageCode   string `json:"his_usage_code,omitempty" xml:"his_usage_code,omitempty"`
    

    // 用法名称(非空)
    
    HisUsageName   string `json:"his_usage_name,omitempty" xml:"his_usage_name,omitempty"`
    

    // 渠道(非空)
    
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
    

    // 1可用0停用(非空)
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

}

package alidoc

// FrequencyRequest 
type FrequencyRequest struct {
    // 用药频次（非空）
    FreqTimes   string `json:"freq_times,omitempty" xml:"freq_times,omitempty"`
    // 状态1：正常 0：停用（非空）
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 频次名称（非空）
    HisFrequencyName   string `json:"his_frequency_name,omitempty" xml:"his_frequency_name,omitempty"`
    // 频次编码（非空）
    HisFrequencyCode   string `json:"his_frequency_code,omitempty" xml:"his_frequency_code,omitempty"`
    // 渠道类型（非空）
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}

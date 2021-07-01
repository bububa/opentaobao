package ascpchannel

// AlibabaAscpChannelDistributorInventoryGetResultDto 
type AlibabaAscpChannelDistributorInventoryGetResultDto struct {
    // 调用链路ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 结果
    Datas   []ChannelInventoryDto `json:"datas,omitempty" xml:"datas>channel_inventory_dto,omitempty"`
}

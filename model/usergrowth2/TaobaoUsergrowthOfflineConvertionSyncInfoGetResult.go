package usergrowth2

// TaobaoUsergrowthOfflineConvertionSyncInfoGetResult 
type TaobaoUsergrowthOfflineConvertionSyncInfoGetResult struct {
    // 跟踪id
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回素材id
    Data   *OfflineConvertionSyncInfoDTO `json:"data,omitempty" xml:"data,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 是否调用成功
    Successful   bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

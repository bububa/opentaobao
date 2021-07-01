package ascpchannel

// AlibabaAscpIndustryLogisticsSyncResultDto 
type AlibabaAscpIndustryLogisticsSyncResultDto struct {
    // 错误描述
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 错误编码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 数据内容
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
}

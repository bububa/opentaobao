package alilabs

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult 结构体
type AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult struct {
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 接口返回model
	Result *DeviceStatusVO `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码（200：成功，其他：失败）
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

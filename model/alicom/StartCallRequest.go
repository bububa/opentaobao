package alicom

// StartCallRequest 结构体
type StartCallRequest struct {
	// AXN分机号产品中通过IVR放音收取上来的用户输入的分机字符
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 中间号码
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 主叫号码
	CallNo string `json:"call_no,omitempty" xml:"call_no,omitempty"`
	// 呼叫开始时间
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 唯一的呼叫ID，最大可支持字符串长度256
	CallId string `json:"call_id,omitempty" xml:"call_id,omitempty"`
	// 行为类型,CALL:呼叫行为,SMS:短信行为
	RecordType string `json:"record_type,omitempty" xml:"record_type,omitempty"`
	// 供应商KEY
	VendorKey string `json:"vendor_key,omitempty" xml:"vendor_key,omitempty"`
}

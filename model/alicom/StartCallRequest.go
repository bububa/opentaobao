package alicom

import (
	"sync"
)

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
	// 呼叫能力阶段，默认填0
	CallPhase string `json:"call_phase,omitempty" xml:"call_phase,omitempty"`
	// 如果有原始被叫号码信息填1
	BCallHistory string `json:"b_call_history,omitempty" xml:"b_call_history,omitempty"`
	// 当响应指令为“短信解析”时，供应商平台完成短信内容解析，在重新发起查询请求时会携带；短信解析结果，0：成功，1：失败
	ParseResult string `json:"parse_result,omitempty" xml:"parse_result,omitempty"`
}

var poolStartCallRequest = sync.Pool{
	New: func() any {
		return new(StartCallRequest)
	},
}

// GetStartCallRequest() 从对象池中获取StartCallRequest
func GetStartCallRequest() *StartCallRequest {
	return poolStartCallRequest.Get().(*StartCallRequest)
}

// ReleaseStartCallRequest 释放StartCallRequest
func ReleaseStartCallRequest(v *StartCallRequest) {
	v.Extension = ""
	v.SecretNo = ""
	v.CallNo = ""
	v.CallTime = ""
	v.CallId = ""
	v.RecordType = ""
	v.VendorKey = ""
	v.CallPhase = ""
	v.BCallHistory = ""
	v.ParseResult = ""
	poolStartCallRequest.Put(v)
}

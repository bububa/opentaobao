package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoshrbackgroundreportnotifyAPIResponse 背调公司背调结果通知 API返回值
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
type AlibabamoshrbackgroundreportnotifyAPIResponse struct {
	model.CommonResponse
	AlibabamoshrbackgroundreportnotifyAPIResponseModel
}

// AlibabamoshrbackgroundreportnotifyAPIResponseModel is 背调公司背调结果通知 成功返回结果
type AlibabamoshrbackgroundreportnotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_hr_background_report_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 异步结果
	AsyncResult string `json:"async_result,omitempty" xml:"async_result,omitempty"`
	// 返回对象
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误code
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 调用是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// 是否同步
	IsAsync bool `json:"is_async,omitempty" xml:"is_async,omitempty"`
}

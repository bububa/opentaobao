package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosHrBackgroundReportNotifyAPIResponse 背调公司背调结果通知 API返回值
// alibaba.mos.hr.background.report.notify
//
// 背调公司背调结果通知
type AlibabaMosHrBackgroundReportNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaMosHrBackgroundReportNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosHrBackgroundReportNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosHrBackgroundReportNotifyAPIResponseModel).Reset()
}

// AlibabaMosHrBackgroundReportNotifyAPIResponseModel is 背调公司背调结果通知 成功返回结果
type AlibabaMosHrBackgroundReportNotifyAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaMosHrBackgroundReportNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.AsyncResult = ""
	m.Data = ""
	m.Errcode = ""
	m.Attributes = ""
	m.ErrMessage = ""
	m.Issuccess = false
	m.IsAsync = false
}

var poolAlibabaMosHrBackgroundReportNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosHrBackgroundReportNotifyAPIResponse)
	},
}

// GetAlibabaMosHrBackgroundReportNotifyAPIResponse 从 sync.Pool 获取 AlibabaMosHrBackgroundReportNotifyAPIResponse
func GetAlibabaMosHrBackgroundReportNotifyAPIResponse() *AlibabaMosHrBackgroundReportNotifyAPIResponse {
	return poolAlibabaMosHrBackgroundReportNotifyAPIResponse.Get().(*AlibabaMosHrBackgroundReportNotifyAPIResponse)
}

// ReleaseAlibabaMosHrBackgroundReportNotifyAPIResponse 将 AlibabaMosHrBackgroundReportNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosHrBackgroundReportNotifyAPIResponse(v *AlibabaMosHrBackgroundReportNotifyAPIResponse) {
	v.Reset()
	poolAlibabaMosHrBackgroundReportNotifyAPIResponse.Put(v)
}

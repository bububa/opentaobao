package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosflowWorkStartprocessAPIResponse 发起流程 API返回值
// alibaba.mosflow.work.startprocess
//
// 业务发起流程审批
type AlibabaMosflowWorkStartprocessAPIResponse struct {
	model.CommonResponse
	AlibabaMosflowWorkStartprocessAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosflowWorkStartprocessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosflowWorkStartprocessAPIResponseModel).Reset()
}

// AlibabaMosflowWorkStartprocessAPIResponseModel is 发起流程 成功返回结果
type AlibabaMosflowWorkStartprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mosflow_work_startprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 异常信息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 异常Code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosflowWorkStartprocessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.ResultMessage = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolAlibabaMosflowWorkStartprocessAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosflowWorkStartprocessAPIResponse)
	},
}

// GetAlibabaMosflowWorkStartprocessAPIResponse 从 sync.Pool 获取 AlibabaMosflowWorkStartprocessAPIResponse
func GetAlibabaMosflowWorkStartprocessAPIResponse() *AlibabaMosflowWorkStartprocessAPIResponse {
	return poolAlibabaMosflowWorkStartprocessAPIResponse.Get().(*AlibabaMosflowWorkStartprocessAPIResponse)
}

// ReleaseAlibabaMosflowWorkStartprocessAPIResponse 将 AlibabaMosflowWorkStartprocessAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosflowWorkStartprocessAPIResponse(v *AlibabaMosflowWorkStartprocessAPIResponse) {
	v.Reset()
	poolAlibabaMosflowWorkStartprocessAPIResponse.Put(v)
}

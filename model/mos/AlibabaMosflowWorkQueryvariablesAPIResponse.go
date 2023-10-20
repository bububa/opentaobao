package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosflowWorkQueryvariablesAPIResponse 获取指定流程上下文参数 API返回值
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
type AlibabaMosflowWorkQueryvariablesAPIResponse struct {
	model.CommonResponse
	AlibabaMosflowWorkQueryvariablesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosflowWorkQueryvariablesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosflowWorkQueryvariablesAPIResponseModel).Reset()
}

// AlibabaMosflowWorkQueryvariablesAPIResponseModel is 获取指定流程上下文参数 成功返回结果
type AlibabaMosflowWorkQueryvariablesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mosflow_work_queryvariables_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MultiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosflowWorkQueryvariablesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosflowWorkQueryvariablesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosflowWorkQueryvariablesAPIResponse)
	},
}

// GetAlibabaMosflowWorkQueryvariablesAPIResponse 从 sync.Pool 获取 AlibabaMosflowWorkQueryvariablesAPIResponse
func GetAlibabaMosflowWorkQueryvariablesAPIResponse() *AlibabaMosflowWorkQueryvariablesAPIResponse {
	return poolAlibabaMosflowWorkQueryvariablesAPIResponse.Get().(*AlibabaMosflowWorkQueryvariablesAPIResponse)
}

// ReleaseAlibabaMosflowWorkQueryvariablesAPIResponse 将 AlibabaMosflowWorkQueryvariablesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosflowWorkQueryvariablesAPIResponse(v *AlibabaMosflowWorkQueryvariablesAPIResponse) {
	v.Reset()
	poolAlibabaMosflowWorkQueryvariablesAPIResponse.Put(v)
}

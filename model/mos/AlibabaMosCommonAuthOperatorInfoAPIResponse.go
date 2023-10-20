package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommonAuthOperatorInfoAPIResponse 获取当前人员信息 API返回值
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
type AlibabaMosCommonAuthOperatorInfoAPIResponse struct {
	model.CommonResponse
	AlibabaMosCommonAuthOperatorInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosCommonAuthOperatorInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosCommonAuthOperatorInfoAPIResponseModel).Reset()
}

// AlibabaMosCommonAuthOperatorInfoAPIResponseModel is 获取当前人员信息 成功返回结果
type AlibabaMosCommonAuthOperatorInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_common_auth_operator_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 异步结果
	AsyncResult string `json:"async_result,omitempty" xml:"async_result,omitempty"`
	// 错误code
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 授权的人员信息
	Data *OperatorUserInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// 是否同步
	IsAsync bool `json:"is_async,omitempty" xml:"is_async,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosCommonAuthOperatorInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.AsyncResult = ""
	m.Errcode = ""
	m.Attributes = ""
	m.ErrMessage = ""
	m.Data = nil
	m.Issuccess = false
	m.IsAsync = false
}

var poolAlibabaMosCommonAuthOperatorInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosCommonAuthOperatorInfoAPIResponse)
	},
}

// GetAlibabaMosCommonAuthOperatorInfoAPIResponse 从 sync.Pool 获取 AlibabaMosCommonAuthOperatorInfoAPIResponse
func GetAlibabaMosCommonAuthOperatorInfoAPIResponse() *AlibabaMosCommonAuthOperatorInfoAPIResponse {
	return poolAlibabaMosCommonAuthOperatorInfoAPIResponse.Get().(*AlibabaMosCommonAuthOperatorInfoAPIResponse)
}

// ReleaseAlibabaMosCommonAuthOperatorInfoAPIResponse 将 AlibabaMosCommonAuthOperatorInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosCommonAuthOperatorInfoAPIResponse(v *AlibabaMosCommonAuthOperatorInfoAPIResponse) {
	v.Reset()
	poolAlibabaMosCommonAuthOperatorInfoAPIResponse.Put(v)
}

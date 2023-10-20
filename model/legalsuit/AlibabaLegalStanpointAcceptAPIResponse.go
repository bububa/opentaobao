package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStanpointAcceptAPIResponse 采纳口径 API返回值
// alibaba.legal.stanpoint.accept
//
// 采纳口径
type AlibabaLegalStanpointAcceptAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStanpointAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStanpointAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStanpointAcceptAPIResponseModel).Reset()
}

// AlibabaLegalStanpointAcceptAPIResponseModel is 采纳口径 成功返回结果
type AlibabaLegalStanpointAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_stanpoint_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态code
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
	// 返回内容
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStanpointAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
	m.Content = false
}

var poolAlibabaLegalStanpointAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStanpointAcceptAPIResponse)
	},
}

// GetAlibabaLegalStanpointAcceptAPIResponse 从 sync.Pool 获取 AlibabaLegalStanpointAcceptAPIResponse
func GetAlibabaLegalStanpointAcceptAPIResponse() *AlibabaLegalStanpointAcceptAPIResponse {
	return poolAlibabaLegalStanpointAcceptAPIResponse.Get().(*AlibabaLegalStanpointAcceptAPIResponse)
}

// ReleaseAlibabaLegalStanpointAcceptAPIResponse 将 AlibabaLegalStanpointAcceptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStanpointAcceptAPIResponse(v *AlibabaLegalStanpointAcceptAPIResponse) {
	v.Reset()
	poolAlibabaLegalStanpointAcceptAPIResponse.Put(v)
}

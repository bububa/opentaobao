package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointDeleteAPIResponse 删除关联口径 API返回值
// alibaba.legal.standpoint.delete
//
// 删除关联口径
type AlibabaLegalStandpointDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointDeleteAPIResponseModel).Reset()
}

// AlibabaLegalStandpointDeleteAPIResponseModel is 删除关联口径 成功返回结果
type AlibabaLegalStandpointDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_delete_response"`
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
func (m *AlibabaLegalStandpointDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
	m.Content = false
}

var poolAlibabaLegalStandpointDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointDeleteAPIResponse)
	},
}

// GetAlibabaLegalStandpointDeleteAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointDeleteAPIResponse
func GetAlibabaLegalStandpointDeleteAPIResponse() *AlibabaLegalStandpointDeleteAPIResponse {
	return poolAlibabaLegalStandpointDeleteAPIResponse.Get().(*AlibabaLegalStandpointDeleteAPIResponse)
}

// ReleaseAlibabaLegalStandpointDeleteAPIResponse 将 AlibabaLegalStandpointDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointDeleteAPIResponse(v *AlibabaLegalStandpointDeleteAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointDeleteAPIResponse.Put(v)
}

package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointQueryAPIResponse 口径查询 API返回值
// alibaba.legal.standpoint.query
//
// 口径查询
type AlibabaLegalStandpointQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointQueryAPIResponseModel).Reset()
}

// AlibabaLegalStandpointQueryAPIResponseModel is 口径查询 成功返回结果
type AlibabaLegalStandpointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回状态
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 分页对象
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = nil
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointQueryAPIResponse)
	},
}

// GetAlibabaLegalStandpointQueryAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointQueryAPIResponse
func GetAlibabaLegalStandpointQueryAPIResponse() *AlibabaLegalStandpointQueryAPIResponse {
	return poolAlibabaLegalStandpointQueryAPIResponse.Get().(*AlibabaLegalStandpointQueryAPIResponse)
}

// ReleaseAlibabaLegalStandpointQueryAPIResponse 将 AlibabaLegalStandpointQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointQueryAPIResponse(v *AlibabaLegalStandpointQueryAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointQueryAPIResponse.Put(v)
}

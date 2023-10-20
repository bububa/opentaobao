package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointStandpointQueryAPIResponse 查询具体口径 API返回值
// alibaba.legal.standpoint.standpoint.query
//
// 查询具体口径
type AlibabaLegalStandpointStandpointQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointStandpointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointStandpointQueryAPIResponseModel).Reset()
}

// AlibabaLegalStandpointStandpointQueryAPIResponseModel is 查询具体口径 成功返回结果
type AlibabaLegalStandpointStandpointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_standpoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 输出对象
	Content *StandpointOutPutDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointStandpointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.Content = nil
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointStandpointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointStandpointQueryAPIResponse)
	},
}

// GetAlibabaLegalStandpointStandpointQueryAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointStandpointQueryAPIResponse
func GetAlibabaLegalStandpointStandpointQueryAPIResponse() *AlibabaLegalStandpointStandpointQueryAPIResponse {
	return poolAlibabaLegalStandpointStandpointQueryAPIResponse.Get().(*AlibabaLegalStandpointStandpointQueryAPIResponse)
}

// ReleaseAlibabaLegalStandpointStandpointQueryAPIResponse 将 AlibabaLegalStandpointStandpointQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointStandpointQueryAPIResponse(v *AlibabaLegalStandpointStandpointQueryAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointStandpointQueryAPIResponse.Put(v)
}

package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseQuerystandpointSaveAPIResponse 法宝侧主动查询反馈 API返回值
// alibaba.legal.case.querystandpoint.save
//
// 法宝侧主动查询反馈口径,此接口只用来反馈主动查询的口径,之前推送的口径反馈不适合
type AlibabaLegalCaseQuerystandpointSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseQuerystandpointSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseQuerystandpointSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseQuerystandpointSaveAPIResponseModel).Reset()
}

// AlibabaLegalCaseQuerystandpointSaveAPIResponseModel is 法宝侧主动查询反馈 成功返回结果
type AlibabaLegalCaseQuerystandpointSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_querystandpoint_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	ErrorNum string `json:"error_num,omitempty" xml:"error_num,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否反馈传成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseQuerystandpointSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorNum = ""
	m.ErrorMsg = ""
	m.Content = false
}

var poolAlibabaLegalCaseQuerystandpointSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseQuerystandpointSaveAPIResponse)
	},
}

// GetAlibabaLegalCaseQuerystandpointSaveAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseQuerystandpointSaveAPIResponse
func GetAlibabaLegalCaseQuerystandpointSaveAPIResponse() *AlibabaLegalCaseQuerystandpointSaveAPIResponse {
	return poolAlibabaLegalCaseQuerystandpointSaveAPIResponse.Get().(*AlibabaLegalCaseQuerystandpointSaveAPIResponse)
}

// ReleaseAlibabaLegalCaseQuerystandpointSaveAPIResponse 将 AlibabaLegalCaseQuerystandpointSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseQuerystandpointSaveAPIResponse(v *AlibabaLegalCaseQuerystandpointSaveAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseQuerystandpointSaveAPIResponse.Put(v)
}

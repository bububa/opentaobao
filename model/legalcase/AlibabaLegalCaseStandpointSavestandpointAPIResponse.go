package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointSavestandpointAPIResponse 新增反馈口径 API返回值
// alibaba.legal.case.standpoint.savestandpoint
//
// 新增反馈口径 ,从外部接受反馈的口径
type AlibabaLegalCaseStandpointSavestandpointAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseStandpointSavestandpointAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointSavestandpointAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseStandpointSavestandpointAPIResponseModel).Reset()
}

// AlibabaLegalCaseStandpointSavestandpointAPIResponseModel is 新增反馈口径 成功返回结果
type AlibabaLegalCaseStandpointSavestandpointAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_savestandpoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 反馈的新增口径id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointSavestandpointAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.Content = 0
	m.IsSuccess = false
}

var poolAlibabaLegalCaseStandpointSavestandpointAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseStandpointSavestandpointAPIResponse)
	},
}

// GetAlibabaLegalCaseStandpointSavestandpointAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseStandpointSavestandpointAPIResponse
func GetAlibabaLegalCaseStandpointSavestandpointAPIResponse() *AlibabaLegalCaseStandpointSavestandpointAPIResponse {
	return poolAlibabaLegalCaseStandpointSavestandpointAPIResponse.Get().(*AlibabaLegalCaseStandpointSavestandpointAPIResponse)
}

// ReleaseAlibabaLegalCaseStandpointSavestandpointAPIResponse 将 AlibabaLegalCaseStandpointSavestandpointAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseStandpointSavestandpointAPIResponse(v *AlibabaLegalCaseStandpointSavestandpointAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseStandpointSavestandpointAPIResponse.Put(v)
}

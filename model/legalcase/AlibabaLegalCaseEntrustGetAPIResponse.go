package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseEntrustGetAPIResponse 委托 API返回值
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
type AlibabaLegalCaseEntrustGetAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseEntrustGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseEntrustGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseEntrustGetAPIResponseModel).Reset()
}

// AlibabaLegalCaseEntrustGetAPIResponseModel is 委托 成功返回结果
type AlibabaLegalCaseEntrustGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_entrust_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseEntrustGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalCaseEntrustGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseEntrustGetAPIResponse)
	},
}

// GetAlibabaLegalCaseEntrustGetAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseEntrustGetAPIResponse
func GetAlibabaLegalCaseEntrustGetAPIResponse() *AlibabaLegalCaseEntrustGetAPIResponse {
	return poolAlibabaLegalCaseEntrustGetAPIResponse.Get().(*AlibabaLegalCaseEntrustGetAPIResponse)
}

// ReleaseAlibabaLegalCaseEntrustGetAPIResponse 将 AlibabaLegalCaseEntrustGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseEntrustGetAPIResponse(v *AlibabaLegalCaseEntrustGetAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseEntrustGetAPIResponse.Put(v)
}

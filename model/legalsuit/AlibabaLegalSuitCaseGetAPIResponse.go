package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCaseGetAPIResponse 获取案件信息接口v2版本 API返回值
// alibaba.legal.suit.case.get
//
// 获取案件信息
type AlibabaLegalSuitCaseGetAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCaseGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCaseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCaseGetAPIResponseModel).Reset()
}

// AlibabaLegalSuitCaseGetAPIResponseModel is 获取案件信息接口v2版本 成功返回结果
type AlibabaLegalSuitCaseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_case_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCaseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCaseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCaseGetAPIResponse)
	},
}

// GetAlibabaLegalSuitCaseGetAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCaseGetAPIResponse
func GetAlibabaLegalSuitCaseGetAPIResponse() *AlibabaLegalSuitCaseGetAPIResponse {
	return poolAlibabaLegalSuitCaseGetAPIResponse.Get().(*AlibabaLegalSuitCaseGetAPIResponse)
}

// ReleaseAlibabaLegalSuitCaseGetAPIResponse 将 AlibabaLegalSuitCaseGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCaseGetAPIResponse(v *AlibabaLegalSuitCaseGetAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCaseGetAPIResponse.Put(v)
}

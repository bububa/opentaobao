package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseCommonEnumdataAPIResponse 获取通用枚举值接口 API返回值
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
type AlibabaLegalCaseCommonEnumdataAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseCommonEnumdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCommonEnumdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseCommonEnumdataAPIResponseModel).Reset()
}

// AlibabaLegalCaseCommonEnumdataAPIResponseModel is 获取通用枚举值接口 成功返回结果
type AlibabaLegalCaseCommonEnumdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_common_enumdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCommonEnumdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalCaseCommonEnumdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseCommonEnumdataAPIResponse)
	},
}

// GetAlibabaLegalCaseCommonEnumdataAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseCommonEnumdataAPIResponse
func GetAlibabaLegalCaseCommonEnumdataAPIResponse() *AlibabaLegalCaseCommonEnumdataAPIResponse {
	return poolAlibabaLegalCaseCommonEnumdataAPIResponse.Get().(*AlibabaLegalCaseCommonEnumdataAPIResponse)
}

// ReleaseAlibabaLegalCaseCommonEnumdataAPIResponse 将 AlibabaLegalCaseCommonEnumdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseCommonEnumdataAPIResponse(v *AlibabaLegalCaseCommonEnumdataAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseCommonEnumdataAPIResponse.Put(v)
}

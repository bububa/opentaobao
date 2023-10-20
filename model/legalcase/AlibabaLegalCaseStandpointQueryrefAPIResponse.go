package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointQueryrefAPIResponse 查询推送口径信息 API返回值
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
type AlibabaLegalCaseStandpointQueryrefAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseStandpointQueryrefAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointQueryrefAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseStandpointQueryrefAPIResponseModel).Reset()
}

// AlibabaLegalCaseStandpointQueryrefAPIResponseModel is 查询推送口径信息 成功返回结果
type AlibabaLegalCaseStandpointQueryrefAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_standpoint_queryref_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseStandpointQueryrefAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalCaseStandpointQueryrefAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseStandpointQueryrefAPIResponse)
	},
}

// GetAlibabaLegalCaseStandpointQueryrefAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseStandpointQueryrefAPIResponse
func GetAlibabaLegalCaseStandpointQueryrefAPIResponse() *AlibabaLegalCaseStandpointQueryrefAPIResponse {
	return poolAlibabaLegalCaseStandpointQueryrefAPIResponse.Get().(*AlibabaLegalCaseStandpointQueryrefAPIResponse)
}

// ReleaseAlibabaLegalCaseStandpointQueryrefAPIResponse 将 AlibabaLegalCaseStandpointQueryrefAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseStandpointQueryrefAPIResponse(v *AlibabaLegalCaseStandpointQueryrefAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseStandpointQueryrefAPIResponse.Put(v)
}

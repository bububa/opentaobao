package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtEntrustGetAPIResponse 委托开庭服务查询 API返回值
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
type AlibabaLegalSuitCourtEntrustGetAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourtEntrustGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtEntrustGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourtEntrustGetAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourtEntrustGetAPIResponseModel is 委托开庭服务查询 成功返回结果
type AlibabaLegalSuitCourtEntrustGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_entrust_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtEntrustGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCourtEntrustGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourtEntrustGetAPIResponse)
	},
}

// GetAlibabaLegalSuitCourtEntrustGetAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourtEntrustGetAPIResponse
func GetAlibabaLegalSuitCourtEntrustGetAPIResponse() *AlibabaLegalSuitCourtEntrustGetAPIResponse {
	return poolAlibabaLegalSuitCourtEntrustGetAPIResponse.Get().(*AlibabaLegalSuitCourtEntrustGetAPIResponse)
}

// ReleaseAlibabaLegalSuitCourtEntrustGetAPIResponse 将 AlibabaLegalSuitCourtEntrustGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourtEntrustGetAPIResponse(v *AlibabaLegalSuitCourtEntrustGetAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourtEntrustGetAPIResponse.Put(v)
}

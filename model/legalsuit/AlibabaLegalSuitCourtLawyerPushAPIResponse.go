package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtLawyerPushAPIResponse 推荐律师接口 API返回值
// alibaba.legal.suit.court.lawyer.push
//
// 为诉讼系统推荐律师
type AlibabaLegalSuitCourtLawyerPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourtLawyerPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtLawyerPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourtLawyerPushAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourtLawyerPushAPIResponseModel is 推荐律师接口 成功返回结果
type AlibabaLegalSuitCourtLawyerPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_lawyer_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtLawyerPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCourtLawyerPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourtLawyerPushAPIResponse)
	},
}

// GetAlibabaLegalSuitCourtLawyerPushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourtLawyerPushAPIResponse
func GetAlibabaLegalSuitCourtLawyerPushAPIResponse() *AlibabaLegalSuitCourtLawyerPushAPIResponse {
	return poolAlibabaLegalSuitCourtLawyerPushAPIResponse.Get().(*AlibabaLegalSuitCourtLawyerPushAPIResponse)
}

// ReleaseAlibabaLegalSuitCourtLawyerPushAPIResponse 将 AlibabaLegalSuitCourtLawyerPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourtLawyerPushAPIResponse(v *AlibabaLegalSuitCourtLawyerPushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourtLawyerPushAPIResponse.Put(v)
}

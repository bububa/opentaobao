package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtOpenPushAPIResponse 开庭信息推送接口 API返回值
// alibaba.legal.suit.court.open.push
//
// 供ISV推送开庭信息给集团诉讼
type AlibabaLegalSuitCourtOpenPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourtOpenPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtOpenPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourtOpenPushAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourtOpenPushAPIResponseModel is 开庭信息推送接口 成功返回结果
type AlibabaLegalSuitCourtOpenPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_open_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtOpenPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCourtOpenPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourtOpenPushAPIResponse)
	},
}

// GetAlibabaLegalSuitCourtOpenPushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourtOpenPushAPIResponse
func GetAlibabaLegalSuitCourtOpenPushAPIResponse() *AlibabaLegalSuitCourtOpenPushAPIResponse {
	return poolAlibabaLegalSuitCourtOpenPushAPIResponse.Get().(*AlibabaLegalSuitCourtOpenPushAPIResponse)
}

// ReleaseAlibabaLegalSuitCourtOpenPushAPIResponse 将 AlibabaLegalSuitCourtOpenPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourtOpenPushAPIResponse(v *AlibabaLegalSuitCourtOpenPushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourtOpenPushAPIResponse.Put(v)
}

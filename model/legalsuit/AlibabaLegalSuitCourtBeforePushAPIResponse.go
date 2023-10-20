package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtBeforePushAPIResponse 更新或保存庭前信息 API返回值
// alibaba.legal.suit.court.before.push
//
// 更新或者保存庭前信息
type AlibabaLegalSuitCourtBeforePushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourtBeforePushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtBeforePushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourtBeforePushAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourtBeforePushAPIResponseModel is 更新或保存庭前信息 成功返回结果
type AlibabaLegalSuitCourtBeforePushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_before_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtBeforePushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCourtBeforePushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourtBeforePushAPIResponse)
	},
}

// GetAlibabaLegalSuitCourtBeforePushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourtBeforePushAPIResponse
func GetAlibabaLegalSuitCourtBeforePushAPIResponse() *AlibabaLegalSuitCourtBeforePushAPIResponse {
	return poolAlibabaLegalSuitCourtBeforePushAPIResponse.Get().(*AlibabaLegalSuitCourtBeforePushAPIResponse)
}

// ReleaseAlibabaLegalSuitCourtBeforePushAPIResponse 将 AlibabaLegalSuitCourtBeforePushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourtBeforePushAPIResponse(v *AlibabaLegalSuitCourtBeforePushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourtBeforePushAPIResponse.Put(v)
}

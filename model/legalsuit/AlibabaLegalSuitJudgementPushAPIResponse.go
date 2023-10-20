package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitJudgementPushAPIResponse 推送裁判登记信息给集团法务系统 API返回值
// alibaba.legal.suit.judgement.push
//
// isv推送裁判登记信息给集团法务系统
type AlibabaLegalSuitJudgementPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitJudgementPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitJudgementPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitJudgementPushAPIResponseModel).Reset()
}

// AlibabaLegalSuitJudgementPushAPIResponseModel is 推送裁判登记信息给集团法务系统 成功返回结果
type AlibabaLegalSuitJudgementPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_judgement_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitJudgementPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitJudgementPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitJudgementPushAPIResponse)
	},
}

// GetAlibabaLegalSuitJudgementPushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitJudgementPushAPIResponse
func GetAlibabaLegalSuitJudgementPushAPIResponse() *AlibabaLegalSuitJudgementPushAPIResponse {
	return poolAlibabaLegalSuitJudgementPushAPIResponse.Get().(*AlibabaLegalSuitJudgementPushAPIResponse)
}

// ReleaseAlibabaLegalSuitJudgementPushAPIResponse 将 AlibabaLegalSuitJudgementPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitJudgementPushAPIResponse(v *AlibabaLegalSuitJudgementPushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitJudgementPushAPIResponse.Put(v)
}

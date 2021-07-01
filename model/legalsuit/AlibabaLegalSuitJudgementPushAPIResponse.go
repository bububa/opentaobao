package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitJudgementPushAPIResponse
推送裁判登记信息给集团法务系统 API返回值
alibaba.legal.suit.judgement.push

isv推送裁判登记信息给集团法务系统 */
type AlibabaLegalSuitJudgementPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitJudgementPushAPIResponseModel
}

// AlibabaLegalSuitJudgementPushAPIResponseModel is 推送裁判登记信息给集团法务系统 成功返回结果
type AlibabaLegalSuitJudgementPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_judgement_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

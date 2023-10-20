package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitDominationPushAPIResponse 更新或者保存管辖信息 API返回值
// alibaba.legal.suit.domination.push
//
// ISV推送管辖信息到诉讼平台
type AlibabaLegalSuitDominationPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitDominationPushAPIResponseModel
}

// AlibabaLegalSuitDominationPushAPIResponseModel is 更新或者保存管辖信息 成功返回结果
type AlibabaLegalSuitDominationPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_domination_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

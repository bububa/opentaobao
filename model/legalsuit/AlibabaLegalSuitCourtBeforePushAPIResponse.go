package legalsuit

import (
	"encoding/xml"

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

// AlibabaLegalSuitCourtBeforePushAPIResponseModel is 更新或保存庭前信息 成功返回结果
type AlibabaLegalSuitCourtBeforePushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_before_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

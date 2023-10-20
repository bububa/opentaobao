package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitSealPushAPIResponse 法宝推送用印 API返回值
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
type AlibabaLegalSuitSealPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitSealPushAPIResponseModel
}

// AlibabaLegalSuitSealPushAPIResponseModel is 法宝推送用印 成功返回结果
type AlibabaLegalSuitSealPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_seal_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回内容
	Content []SealResponseModel `json:"content,omitempty" xml:"content>seal_response_model,omitempty"`
	// 错误code
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

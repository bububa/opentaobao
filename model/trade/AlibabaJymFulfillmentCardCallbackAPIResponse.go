package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymfulfillmentcardcallbackAPIResponse 外部商家卡密结果回调 API返回值
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
type AlibabajymfulfillmentcardcallbackAPIResponse struct {
	model.CommonResponse
	AlibabajymfulfillmentcardcallbackAPIResponseModel
}

// AlibabajymfulfillmentcardcallbackAPIResponseModel is 外部商家卡密结果回调 成功返回结果
type AlibabajymfulfillmentcardcallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_fulfillment_card_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败错误码
	FailedCode string `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
	// 失败原因描述
	FailedReason string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
	// 交易猫订单是否成功，true-成功，false-失败
	JymOrderSuccess string `json:"jym_order_success,omitempty" xml:"jym_order_success,omitempty"`
}

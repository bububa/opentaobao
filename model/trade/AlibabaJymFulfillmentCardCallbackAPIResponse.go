package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymFulfillmentCardCallbackAPIResponse 外部商家卡密结果回调 API返回值
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
type AlibabaJymFulfillmentCardCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaJymFulfillmentCardCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymFulfillmentCardCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymFulfillmentCardCallbackAPIResponseModel).Reset()
}

// AlibabaJymFulfillmentCardCallbackAPIResponseModel is 外部商家卡密结果回调 成功返回结果
type AlibabaJymFulfillmentCardCallbackAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaJymFulfillmentCardCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailedCode = ""
	m.FailedReason = ""
	m.JymOrderSuccess = ""
}

var poolAlibabaJymFulfillmentCardCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymFulfillmentCardCallbackAPIResponse)
	},
}

// GetAlibabaJymFulfillmentCardCallbackAPIResponse 从 sync.Pool 获取 AlibabaJymFulfillmentCardCallbackAPIResponse
func GetAlibabaJymFulfillmentCardCallbackAPIResponse() *AlibabaJymFulfillmentCardCallbackAPIResponse {
	return poolAlibabaJymFulfillmentCardCallbackAPIResponse.Get().(*AlibabaJymFulfillmentCardCallbackAPIResponse)
}

// ReleaseAlibabaJymFulfillmentCardCallbackAPIResponse 将 AlibabaJymFulfillmentCardCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymFulfillmentCardCallbackAPIResponse(v *AlibabaJymFulfillmentCardCallbackAPIResponse) {
	v.Reset()
	poolAlibabaJymFulfillmentCardCallbackAPIResponse.Put(v)
}

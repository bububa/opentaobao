package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymFulfillmentIoschargeCallbackAPIResponse 代充充值回调 API返回值
// alibaba.jym.fulfillment.ioscharge.callback
//
// 代充充值回调
type AlibabaJymFulfillmentIoschargeCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaJymFulfillmentIoschargeCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymFulfillmentIoschargeCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymFulfillmentIoschargeCallbackAPIResponseModel).Reset()
}

// AlibabaJymFulfillmentIoschargeCallbackAPIResponseModel is 代充充值回调 成功返回结果
type AlibabaJymFulfillmentIoschargeCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_fulfillment_ioscharge_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败错误码
	FailedCode string `json:"failed_code,omitempty" xml:"failed_code,omitempty"`
	// 交易猫订单是否成功，true-成功，false-失败
	JymOrderSuccess string `json:"jym_order_success,omitempty" xml:"jym_order_success,omitempty"`
	// 失败原因描述
	FailedReason string `json:"failed_reason,omitempty" xml:"failed_reason,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymFulfillmentIoschargeCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailedCode = ""
	m.JymOrderSuccess = ""
	m.FailedReason = ""
}

var poolAlibabaJymFulfillmentIoschargeCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymFulfillmentIoschargeCallbackAPIResponse)
	},
}

// GetAlibabaJymFulfillmentIoschargeCallbackAPIResponse 从 sync.Pool 获取 AlibabaJymFulfillmentIoschargeCallbackAPIResponse
func GetAlibabaJymFulfillmentIoschargeCallbackAPIResponse() *AlibabaJymFulfillmentIoschargeCallbackAPIResponse {
	return poolAlibabaJymFulfillmentIoschargeCallbackAPIResponse.Get().(*AlibabaJymFulfillmentIoschargeCallbackAPIResponse)
}

// ReleaseAlibabaJymFulfillmentIoschargeCallbackAPIResponse 将 AlibabaJymFulfillmentIoschargeCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymFulfillmentIoschargeCallbackAPIResponse(v *AlibabaJymFulfillmentIoschargeCallbackAPIResponse) {
	v.Reset()
	poolAlibabaJymFulfillmentIoschargeCallbackAPIResponse.Put(v)
}

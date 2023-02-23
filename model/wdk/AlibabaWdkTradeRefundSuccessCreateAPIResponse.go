package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundSuccessCreateAPIResponse 五道口终态逆向订单创建 API返回值
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
type AlibabaWdkTradeRefundSuccessCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeRefundSuccessCreateAPIResponseModel
}

// AlibabaWdkTradeRefundSuccessCreateAPIResponseModel is 五道口终态逆向订单创建 成功返回结果
type AlibabaWdkTradeRefundSuccessCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_refund_success_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单返回结果
	RefundOrderResult *AlibabaWdkTradeRefundSuccessCreateApiResult `json:"refund_order_result,omitempty" xml:"refund_order_result,omitempty"`
}

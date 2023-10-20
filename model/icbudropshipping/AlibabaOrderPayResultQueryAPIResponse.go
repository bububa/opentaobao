package icbudropshipping

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaorderpayresultqueryAPIResponse alibaba查询订单支付结果 API返回值
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
type AlibabaorderpayresultqueryAPIResponse struct {
	model.CommonResponse
	AlibabaorderpayresultqueryAPIResponseModel
}

// AlibabaorderpayresultqueryAPIResponseModel is alibaba查询订单支付结果 成功返回结果
type AlibabaorderpayresultqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_order_pay_result_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pay response
	Value *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`
}

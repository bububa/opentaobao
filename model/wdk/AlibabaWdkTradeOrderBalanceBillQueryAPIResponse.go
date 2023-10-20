package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeorderbalancebillqueryAPIResponse 分页拉取订单数据 API返回值
// alibaba.wdk.trade.order.balance.bill.query
//
// 提供接口供外部调用，分页拉取订单数据
type AlibabawdktradeorderbalancebillqueryAPIResponse struct {
	model.CommonResponse
	AlibabawdktradeorderbalancebillqueryAPIResponseModel
}

// AlibabawdktradeorderbalancebillqueryAPIResponseModel is 分页拉取订单数据 成功返回结果
type AlibabawdktradeorderbalancebillqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_balance_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ApiResult
	ApiResult *AlibabawdktradeorderbalancebillqueryApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderQueryAPIResponse 查询外部交易订单接口 API返回值
// alibaba.wdk.trade.order.query
//
// 通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeOrderQueryAPIResponseModel
}

// AlibabaWdkTradeOrderQueryAPIResponseModel is 查询外部交易订单接口 成功返回结果
type AlibabaWdkTradeOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *TradeOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

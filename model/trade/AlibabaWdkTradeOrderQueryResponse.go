package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部交易订单接口 APIResponse
alibaba.wdk.trade.order.query

通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeOrderQueryResponse
}

type AlibabaWdkTradeOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果
    
    Result   *TradeOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单创单接口 APIResponse
alibaba.wdk.trade.order.create

通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkTradeOrderCreateResponse
}

type AlibabaWdkTradeOrderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_trade_order_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 执行结果
    
    Result   *OrderResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

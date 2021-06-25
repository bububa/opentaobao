package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单创单接口 APIResponse
alibaba.wdk.trade.order.create

通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkTradeOrderCreateResponse `json:"alibaba_wdk_trade_order_create_response,omitempty"`
}

type AlibabaWdkTradeOrderCreateResponse struct {

    // 执行结果
    Result   *OrderResult `json:"result,omitempty"`

}

package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家回调接口 APIResponse
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
*/
type TaobaoWtTradeOrderResultcallbackAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWtTradeOrderResultcallbackResponse `json:"taobao_wt_trade_order_resultcallback_response,omitempty"`
}

type TaobaoWtTradeOrderResultcallbackResponse struct {

    // result
    Result   *CommonRtnDo `json:"result,omitempty"`

}

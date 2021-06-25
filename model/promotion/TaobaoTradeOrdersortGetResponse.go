package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取前N有礼活动的开奖订单列表 APIResponse
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
type TaobaoTradeOrdersortGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTradeOrdersortGetResponse `json:"taobao_trade_ordersort_get_response,omitempty"`
}

type TaobaoTradeOrdersortGetResponse struct {

    // 接口返回model
    Result   *TaobaoTradeOrdersortGetResult `json:"result,omitempty"`

}

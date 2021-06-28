package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
获取前N有礼活动的开奖订单列表 
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表
*/
func TaobaoTradeOrdersortGet(clt *core.SDKClient, req *promotion.TaobaoTradeOrdersortGetRequest, session string) (*promotion.TaobaoTradeOrdersortGetAPIResponse, error) {
    var resp promotion.TaobaoTradeOrdersortGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

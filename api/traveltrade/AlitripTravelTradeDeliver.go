package traveltrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traveltrade"
)

/* 
飞猪度假-订单发货接口 
alitrip.travel.trade.deliver

航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
*/
func AlitripTravelTradeDeliver(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeDeliverRequest, session string) (*traveltrade.AlitripTravelTradeDeliverAPIResponse, error) {
    var resp traveltrade.AlitripTravelTradeDeliverAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

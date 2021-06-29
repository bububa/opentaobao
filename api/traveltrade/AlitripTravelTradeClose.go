package traveltrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traveltrade"
)

/* 
飞猪度假-订单关闭接口（快速退款） 
alitrip.travel.trade.close

卖家关单（快速退款接口），不支持二次预约商品的订单
*/
func AlitripTravelTradeClose(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeCloseRequest, session string) (*traveltrade.AlitripTravelTradeCloseAPIResponse, error) {
    var resp traveltrade.AlitripTravelTradeCloseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

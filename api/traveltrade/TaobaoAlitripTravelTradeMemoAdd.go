package traveltrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traveltrade"
)

/* 
添加一笔交易备注 
taobao.alitrip.travel.trade.memo.add

对一笔交易添加备注
*/
func TaobaoAlitripTravelTradeMemoAdd(clt *core.SDKClient, req *traveltrade.TaobaoAlitripTravelTradeMemoAddAPIRequest, session string) (*traveltrade.TaobaoAlitripTravelTradeMemoAddAPIResponse, error) {
    var resp traveltrade.TaobaoAlitripTravelTradeMemoAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

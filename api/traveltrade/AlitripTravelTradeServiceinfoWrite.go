package traveltrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/traveltrade"
)

/* 
订单服务信息写入接口 
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口
*/
func AlitripTravelTradeServiceinfoWrite(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeServiceinfoWriteRequest, session string) (*traveltrade.AlitripTravelTradeServiceinfoWriteAPIResponse, error) {
    var resp traveltrade.AlitripTravelTradeServiceinfoWriteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
汽车票保险推荐 
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容
*/
func TaobaoAlitripBusTicketsInsuranceRecommend(clt *core.SDKClient, req *bus.TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest, session string) (*bus.TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse, error) {
    var resp bus.TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

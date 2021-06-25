package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
飞猪订单状态查询接口 
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单
*/
func TaobaoAlitripCarOrderQuery(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderQueryRequest, session string) (*car.TaobaoAlitripCarOrderQueryResponse, error) {
    var resp car.TaobaoAlitripCarOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

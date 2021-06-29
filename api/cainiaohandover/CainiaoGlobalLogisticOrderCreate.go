package cainiaohandover

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaohandover"
)

/* 
创建物流订单 
cainiao.global.logistic.order.create

创建物流订单
*/
func CainiaoGlobalLogisticOrderCreate(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalLogisticOrderCreateRequest, session string) (*cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIResponse, error) {
    var resp cainiaohandover.CainiaoGlobalLogisticOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
天猫开新车查询订单id 
tmall.car.lease.orderid.get

天猫开新车查询订单id
*/
func TmallCarLeaseOrderidGet(clt *core.SDKClient, req *tmallcar.TmallCarLeaseOrderidGetRequest, session string) (*tmallcar.TmallCarLeaseOrderidGetAPIResponse, error) {
    var resp tmallcar.TmallCarLeaseOrderidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
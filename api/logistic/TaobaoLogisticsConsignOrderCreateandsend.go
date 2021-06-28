package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
创建订单并发货 
taobao.logistics.consign.order.createandsend

创建物流订单，并发货。
*/
func TaobaoLogisticsConsignOrderCreateandsend(clt *core.SDKClient, req *logistic.TaobaoLogisticsConsignOrderCreateandsendRequest, session string) (*logistic.TaobaoLogisticsConsignOrderCreateandsendAPIResponse, error) {
    var resp logistic.TaobaoLogisticsConsignOrderCreateandsendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
无需物流（虚拟）发货处理 
taobao.logistics.dummy.send

用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
*/
func TaobaoLogisticsDummySend(clt *core.SDKClient, req *logistic.TaobaoLogisticsDummySendRequest, session string) (*logistic.TaobaoLogisticsDummySendAPIResponse, error) {
    var resp logistic.TaobaoLogisticsDummySendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

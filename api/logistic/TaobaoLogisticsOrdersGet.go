package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
批量查询物流订单 
taobao.logistics.orders.get

批量查询物流订单。
*/
func TaobaoLogisticsOrdersGet(clt *core.SDKClient, req *logistic.TaobaoLogisticsOrdersGetRequest, session string) (*logistic.TaobaoLogisticsOrdersGetAPIResponse, error) {
    var resp logistic.TaobaoLogisticsOrdersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

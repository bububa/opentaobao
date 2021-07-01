package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
查询卖家地址库 
taobao.logistics.address.search

通过此接口查询卖家地址库，
*/
func TaobaoLogisticsAddressSearch(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressSearchAPIRequest, session string) (*logistic.TaobaoLogisticsAddressSearchAPIResponse, error) {
    var resp logistic.TaobaoLogisticsAddressSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

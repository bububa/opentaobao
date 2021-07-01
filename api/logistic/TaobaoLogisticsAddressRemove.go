package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
删除卖家地址库 
taobao.logistics.address.remove

用此接口删除卖家地址库
*/
func TaobaoLogisticsAddressRemove(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressRemoveAPIRequest, session string) (*logistic.TaobaoLogisticsAddressRemoveAPIResponse, error) {
    var resp logistic.TaobaoLogisticsAddressRemoveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
卖家地址库修改 
taobao.logistics.address.modify

卖家地址库修改
*/
func TaobaoLogisticsAddressModify(clt *core.SDKClient, req *logistic.TaobaoLogisticsAddressModifyRequest, session string) (*logistic.TaobaoLogisticsAddressModifyResponse, error) {
    var resp logistic.TaobaoLogisticsAddressModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
资金流水查询 
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询
*/
func AlibabaAlicomOrderPreauthorizeQueryFund(clt *core.SDKClient, req *alicom.AlibabaAlicomOrderPreauthorizeQueryFundRequest, session string) (*alicom.AlibabaAlicomOrderPreauthorizeQueryFundResponse, error) {
    var resp alicom.AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

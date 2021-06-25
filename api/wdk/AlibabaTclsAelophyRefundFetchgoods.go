package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
saas 售后逆向 商户发起逆向取货 
alibaba.tcls.aelophy.refund.fetchgoods

saas 售后逆向 商户发起逆向取货
*/
func AlibabaTclsAelophyRefundFetchgoods(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundFetchgoodsRequest, session string) (*wdk.AlibabaTclsAelophyRefundFetchgoodsResponse, error) {
    var resp wdk.AlibabaTclsAelophyRefundFetchgoodsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

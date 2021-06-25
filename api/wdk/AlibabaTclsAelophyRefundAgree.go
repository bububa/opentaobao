package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
saas 售后逆向 商户同意用户逆向申请 
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请
*/
func AlibabaTclsAelophyRefundAgree(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyRefundAgreeRequest, session string) (*wdk.AlibabaTclsAelophyRefundAgreeResponse, error) {
    var resp wdk.AlibabaTclsAelophyRefundAgreeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

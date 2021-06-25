package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
流量平台用户签约情况查询 
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询
*/
func AlibabaAliqinFlowWalletSign(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletSignRequest, session string) (*alicom.AlibabaAliqinFlowWalletSignResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletSignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
查询流量充值状态 
alibaba.aliqin.flow.wallet.query.charge

查询流量充值状态
*/
func AlibabaAliqinFlowWalletQueryCharge(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletQueryChargeRequest, session string) (*alicom.AlibabaAliqinFlowWalletQueryChargeAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletQueryChargeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
流量扣减 
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口
*/
func AlibabaAliqinFlowWalletConsume(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletConsumeRequest, session string) (*alicom.AlibabaAliqinFlowWalletConsumeAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
流量发放 
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发
*/
func AlibabaAliqinFlowWalletSendFlow(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletSendFlowRequest, session string) (*alicom.AlibabaAliqinFlowWalletSendFlowAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletSendFlowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

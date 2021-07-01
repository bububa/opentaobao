package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
流量直充 
alibaba.aliqin.flow.wallet.charge

流量直充
*/
func AlibabaAliqinFlowWalletCharge(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletChargeAPIRequest, session string) (*alicom.AlibabaAliqinFlowWalletChargeAPIResponse, error) {
    var resp alicom.AlibabaAliqinFlowWalletChargeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

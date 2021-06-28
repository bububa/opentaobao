package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量直充 APIResponse
alibaba.aliqin.flow.wallet.charge

流量直充
*/
type AlibabaAliqinFlowWalletChargeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFlowWalletChargeResponse `json:"alibaba_aliqin_flow_wallet_charge_response,omitempty"` 
    AlibabaAliqinFlowWalletChargeResponse
}

/* model for simplify = false
type AlibabaAliqinFlowWalletChargeResponse struct {

    // 充值请求
    
    Charge   string `json:"charge,omitempty"`
    

}
*/

type AlibabaAliqinFlowWalletChargeResponse struct {

    // 充值请求
    Charge   string `json:"charge,omitempty"`

}

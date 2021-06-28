package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询流量充值状态 APIResponse
alibaba.aliqin.flow.wallet.query.charge

查询流量充值状态
*/
type AlibabaAliqinFlowWalletQueryChargeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFlowWalletQueryChargeResponse `json:"alibaba_aliqin_flow_wallet_query_charge_response,omitempty"` 
    AlibabaAliqinFlowWalletQueryChargeResponse
}

/* model for simplify = false
type AlibabaAliqinFlowWalletQueryChargeResponse struct {

    // 充值状态
    
    Charge   string `json:"charge,omitempty"`
    

}
*/

type AlibabaAliqinFlowWalletQueryChargeResponse struct {

    // 充值状态
    Charge   string `json:"charge,omitempty"`

}

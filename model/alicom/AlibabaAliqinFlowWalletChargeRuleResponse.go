package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包直充（根据号码归属地省份路由） APIResponse
alibaba.aliqin.flow.wallet.charge.rule

流量钱包直充（根据号码归属地省份路由）
*/
type AlibabaAliqinFlowWalletChargeRuleAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFlowWalletChargeRuleResponse `json:"alibaba_aliqin_flow_wallet_charge_rule_response,omitempty"` 
    AlibabaAliqinFlowWalletChargeRuleResponse
}

/* model for simplify = false
type AlibabaAliqinFlowWalletChargeRuleResponse struct {

    // {"error":"true","msg":"返回信息"}
    
    Charge   string `json:"charge,omitempty"`
    

}
*/

type AlibabaAliqinFlowWalletChargeRuleResponse struct {

    // {"error":"true","msg":"返回信息"}
    Charge   string `json:"charge,omitempty"`

}

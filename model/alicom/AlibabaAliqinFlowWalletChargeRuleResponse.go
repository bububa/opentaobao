package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包直充（根据号码归属地省份路由） APIResponse
alibaba.aliqin.flow.wallet.charge.rule

流量钱包直充（根据号码归属地省份路由）
*/
type AlibabaAliqinFlowWalletChargeRuleAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_wallet_charge_rule_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // {"error":"true","msg":"返回信息"}
    
    Charge   string `json:"charge,omitempty" xml:"
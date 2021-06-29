package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包直充（根据号码归属地省份路由） API返回值 
alibaba.aliqin.flow.wallet.charge.rule

流量钱包直充（根据号码归属地省份路由）
*/
type AlibabaAliqinFlowWalletChargeRuleAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletChargeRuleResponse
}

// 流量钱包直充（根据号码归属地省份路由） 成功返回结果
type AlibabaAliqinFlowWalletChargeRuleResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_charge_rule_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // {"error":"true","msg":"返回信息"}
    Charge   string `json:"charge,omitempty" xml:"charge,omitempty"`
}

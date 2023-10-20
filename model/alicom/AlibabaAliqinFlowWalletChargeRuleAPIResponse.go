package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletChargeRuleAPIResponse 流量钱包直充（根据号码归属地省份路由） API返回值
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
type AlibabaAliqinFlowWalletChargeRuleAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletChargeRuleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletChargeRuleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletChargeRuleAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletChargeRuleAPIResponseModel is 流量钱包直充（根据号码归属地省份路由） 成功返回结果
type AlibabaAliqinFlowWalletChargeRuleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_charge_rule_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// {&#34;error&#34;:&#34;true&#34;,&#34;msg&#34;:&#34;返回信息&#34;}
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletChargeRuleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Charge = ""
}

var poolAlibabaAliqinFlowWalletChargeRuleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletChargeRuleAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletChargeRuleAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletChargeRuleAPIResponse
func GetAlibabaAliqinFlowWalletChargeRuleAPIResponse() *AlibabaAliqinFlowWalletChargeRuleAPIResponse {
	return poolAlibabaAliqinFlowWalletChargeRuleAPIResponse.Get().(*AlibabaAliqinFlowWalletChargeRuleAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletChargeRuleAPIResponse 将 AlibabaAliqinFlowWalletChargeRuleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletChargeRuleAPIResponse(v *AlibabaAliqinFlowWalletChargeRuleAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletChargeRuleAPIResponse.Put(v)
}

package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletQueryChargeAPIResponse 查询流量充值状态 API返回值
// alibaba.aliqin.flow.wallet.query.charge
//
// 查询流量充值状态
type AlibabaAliqinFlowWalletQueryChargeAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletQueryChargeAPIResponseModel
}

// AlibabaAliqinFlowWalletQueryChargeAPIResponseModel is 查询流量充值状态 成功返回结果
type AlibabaAliqinFlowWalletQueryChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_query_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 充值状态
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}

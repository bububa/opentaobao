package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletChargeAPIResponse 流量直充 API返回值
// alibaba.aliqin.flow.wallet.charge
//
// 流量直充
type AlibabaAliqinFlowWalletChargeAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletChargeAPIResponseModel
}

// AlibabaAliqinFlowWalletChargeAPIResponseModel is 流量直充 成功返回结果
type AlibabaAliqinFlowWalletChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 充值请求
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}

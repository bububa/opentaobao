package alicom

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletChargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletChargeAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletChargeAPIResponseModel is 流量直充 成功返回结果
type AlibabaAliqinFlowWalletChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 充值请求
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletChargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Charge = ""
}

var poolAlibabaAliqinFlowWalletChargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletChargeAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletChargeAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletChargeAPIResponse
func GetAlibabaAliqinFlowWalletChargeAPIResponse() *AlibabaAliqinFlowWalletChargeAPIResponse {
	return poolAlibabaAliqinFlowWalletChargeAPIResponse.Get().(*AlibabaAliqinFlowWalletChargeAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletChargeAPIResponse 将 AlibabaAliqinFlowWalletChargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletChargeAPIResponse(v *AlibabaAliqinFlowWalletChargeAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletChargeAPIResponse.Put(v)
}

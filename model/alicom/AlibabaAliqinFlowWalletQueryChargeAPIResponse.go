package alicom

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletQueryChargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletQueryChargeAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletQueryChargeAPIResponseModel is 查询流量充值状态 成功返回结果
type AlibabaAliqinFlowWalletQueryChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_query_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 充值状态
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletQueryChargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Charge = ""
}

var poolAlibabaAliqinFlowWalletQueryChargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletQueryChargeAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletQueryChargeAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletQueryChargeAPIResponse
func GetAlibabaAliqinFlowWalletQueryChargeAPIResponse() *AlibabaAliqinFlowWalletQueryChargeAPIResponse {
	return poolAlibabaAliqinFlowWalletQueryChargeAPIResponse.Get().(*AlibabaAliqinFlowWalletQueryChargeAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletQueryChargeAPIResponse 将 AlibabaAliqinFlowWalletQueryChargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletQueryChargeAPIResponse(v *AlibabaAliqinFlowWalletQueryChargeAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletQueryChargeAPIResponse.Put(v)
}

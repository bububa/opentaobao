package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletConsumeAPIResponse 流量扣减 API返回值
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
type AlibabaAliqinFlowWalletConsumeAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletConsumeAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletConsumeAPIResponseModel is 流量扣减 成功返回结果
type AlibabaAliqinFlowWalletConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true为成功
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
}

var poolAlibabaAliqinFlowWalletConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletConsumeAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletConsumeAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletConsumeAPIResponse
func GetAlibabaAliqinFlowWalletConsumeAPIResponse() *AlibabaAliqinFlowWalletConsumeAPIResponse {
	return poolAlibabaAliqinFlowWalletConsumeAPIResponse.Get().(*AlibabaAliqinFlowWalletConsumeAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletConsumeAPIResponse 将 AlibabaAliqinFlowWalletConsumeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletConsumeAPIResponse(v *AlibabaAliqinFlowWalletConsumeAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletConsumeAPIResponse.Put(v)
}

package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletCheckBalanceAPIResponse 商家预存余额检查 API返回值
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
type AlibabaAliqinFlowWalletCheckBalanceAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowWalletCheckBalanceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletCheckBalanceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFlowWalletCheckBalanceAPIResponseModel).Reset()
}

// AlibabaAliqinFlowWalletCheckBalanceAPIResponseModel is 商家预存余额检查 成功返回结果
type AlibabaAliqinFlowWalletCheckBalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_check_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 余额是否大于校验值，大于返回true，小于返回false
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFlowWalletCheckBalanceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = ""
}

var poolAlibabaAliqinFlowWalletCheckBalanceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFlowWalletCheckBalanceAPIResponse)
	},
}

// GetAlibabaAliqinFlowWalletCheckBalanceAPIResponse 从 sync.Pool 获取 AlibabaAliqinFlowWalletCheckBalanceAPIResponse
func GetAlibabaAliqinFlowWalletCheckBalanceAPIResponse() *AlibabaAliqinFlowWalletCheckBalanceAPIResponse {
	return poolAlibabaAliqinFlowWalletCheckBalanceAPIResponse.Get().(*AlibabaAliqinFlowWalletCheckBalanceAPIResponse)
}

// ReleaseAlibabaAliqinFlowWalletCheckBalanceAPIResponse 将 AlibabaAliqinFlowWalletCheckBalanceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFlowWalletCheckBalanceAPIResponse(v *AlibabaAliqinFlowWalletCheckBalanceAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFlowWalletCheckBalanceAPIResponse.Put(v)
}

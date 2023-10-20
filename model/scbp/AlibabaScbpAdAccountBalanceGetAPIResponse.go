package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdAccountBalanceGetAPIResponse 查询账户余额 API返回值
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
type AlibabaScbpAdAccountBalanceGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdAccountBalanceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdAccountBalanceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdAccountBalanceGetAPIResponseModel).Reset()
}

// AlibabaScbpAdAccountBalanceGetAPIResponseModel is 查询账户余额 成功返回结果
type AlibabaScbpAdAccountBalanceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_account_balance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Balance string `json:"balance,omitempty" xml:"balance,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdAccountBalanceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Balance = ""
}

var poolAlibabaScbpAdAccountBalanceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdAccountBalanceGetAPIResponse)
	},
}

// GetAlibabaScbpAdAccountBalanceGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAdAccountBalanceGetAPIResponse
func GetAlibabaScbpAdAccountBalanceGetAPIResponse() *AlibabaScbpAdAccountBalanceGetAPIResponse {
	return poolAlibabaScbpAdAccountBalanceGetAPIResponse.Get().(*AlibabaScbpAdAccountBalanceGetAPIResponse)
}

// ReleaseAlibabaScbpAdAccountBalanceGetAPIResponse 将 AlibabaScbpAdAccountBalanceGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdAccountBalanceGetAPIResponse(v *AlibabaScbpAdAccountBalanceGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdAccountBalanceGetAPIResponse.Put(v)
}

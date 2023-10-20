package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiProviderAccountBalanceAPIResponse 供应商渠道余额 API返回值
// alibaba.happytrip.taxi.provider.account.balance
//
// 查询不同供应商不同渠道账户余额
type AlibabaHappytripTaxiProviderAccountBalanceAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiProviderAccountBalanceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiProviderAccountBalanceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiProviderAccountBalanceAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiProviderAccountBalanceAPIResponseModel is 供应商渠道余额 成功返回结果
type AlibabaHappytripTaxiProviderAccountBalanceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_provider_account_balance_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 数据
	Data *AlibabaHappytripTaxiProviderAccountBalanceData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiProviderAccountBalanceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiProviderAccountBalanceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiProviderAccountBalanceAPIResponse)
	},
}

// GetAlibabaHappytripTaxiProviderAccountBalanceAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiProviderAccountBalanceAPIResponse
func GetAlibabaHappytripTaxiProviderAccountBalanceAPIResponse() *AlibabaHappytripTaxiProviderAccountBalanceAPIResponse {
	return poolAlibabaHappytripTaxiProviderAccountBalanceAPIResponse.Get().(*AlibabaHappytripTaxiProviderAccountBalanceAPIResponse)
}

// ReleaseAlibabaHappytripTaxiProviderAccountBalanceAPIResponse 将 AlibabaHappytripTaxiProviderAccountBalanceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiProviderAccountBalanceAPIResponse(v *AlibabaHappytripTaxiProviderAccountBalanceAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiProviderAccountBalanceAPIResponse.Put(v)
}

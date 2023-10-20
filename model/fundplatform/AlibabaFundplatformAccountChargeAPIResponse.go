package fundplatform

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformAccountChargeAPIResponse 资金平台余额账户充值 API返回值
// alibaba.fundplatform.account.charge
//
// 资金平台余额账户充值【创建账户&amp;返回付款URL】
type AlibabaFundplatformAccountChargeAPIResponse struct {
	model.CommonResponse
	AlibabaFundplatformAccountChargeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountChargeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFundplatformAccountChargeAPIResponseModel).Reset()
}

// AlibabaFundplatformAccountChargeAPIResponseModel is 资金平台余额账户充值 成功返回结果
type AlibabaFundplatformAccountChargeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fundplatform_account_charge_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFundplatformAccountChargeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFundplatformAccountChargeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFundplatformAccountChargeAPIResponse)
	},
}

// GetAlibabaFundplatformAccountChargeAPIResponse 从 sync.Pool 获取 AlibabaFundplatformAccountChargeAPIResponse
func GetAlibabaFundplatformAccountChargeAPIResponse() *AlibabaFundplatformAccountChargeAPIResponse {
	return poolAlibabaFundplatformAccountChargeAPIResponse.Get().(*AlibabaFundplatformAccountChargeAPIResponse)
}

// ReleaseAlibabaFundplatformAccountChargeAPIResponse 将 AlibabaFundplatformAccountChargeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFundplatformAccountChargeAPIResponse(v *AlibabaFundplatformAccountChargeAPIResponse) {
	v.Reset()
	poolAlibabaFundplatformAccountChargeAPIResponse.Put(v)
}

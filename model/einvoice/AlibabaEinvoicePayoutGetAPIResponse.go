package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePayoutGetAPIResponse 获取赔付计时列表数据 API返回值
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
type AlibabaEinvoicePayoutGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePayoutGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePayoutGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePayoutGetAPIResponseModel).Reset()
}

// AlibabaEinvoicePayoutGetAPIResponseModel is 获取赔付计时列表数据 成功返回结果
type AlibabaEinvoicePayoutGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_payout_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OrderRightsResult `json:"result,omitempty" xml:"result,omitempty"`
	// 查询结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePayoutGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
	m.IsSuccess = false
}

var poolAlibabaEinvoicePayoutGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePayoutGetAPIResponse)
	},
}

// GetAlibabaEinvoicePayoutGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePayoutGetAPIResponse
func GetAlibabaEinvoicePayoutGetAPIResponse() *AlibabaEinvoicePayoutGetAPIResponse {
	return poolAlibabaEinvoicePayoutGetAPIResponse.Get().(*AlibabaEinvoicePayoutGetAPIResponse)
}

// ReleaseAlibabaEinvoicePayoutGetAPIResponse 将 AlibabaEinvoicePayoutGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePayoutGetAPIResponse(v *AlibabaEinvoicePayoutGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePayoutGetAPIResponse.Put(v)
}

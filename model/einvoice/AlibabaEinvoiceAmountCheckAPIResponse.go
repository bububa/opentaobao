package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceAmountCheckAPIResponse 开票量核对接口 API返回值
// alibaba.einvoice.amount.check
//
// 跟开票服务商核对历史开票量，用来对账
type AlibabaEinvoiceAmountCheckAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceAmountCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceAmountCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceAmountCheckAPIResponseModel).Reset()
}

// AlibabaEinvoiceAmountCheckAPIResponseModel is 开票量核对接口 成功返回结果
type AlibabaEinvoiceAmountCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_amount_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票量查询结果
	EinvoiceAmountCheckResultList []EinvoiceAmountCheckResult `json:"einvoice_amount_check_result_list,omitempty" xml:"einvoice_amount_check_result_list>einvoice_amount_check_result,omitempty"`
	// 查询结果的数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceAmountCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EinvoiceAmountCheckResultList = m.EinvoiceAmountCheckResultList[:0]
	m.TotalCount = 0
	m.IsSuccess = false
}

var poolAlibabaEinvoiceAmountCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceAmountCheckAPIResponse)
	},
}

// GetAlibabaEinvoiceAmountCheckAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceAmountCheckAPIResponse
func GetAlibabaEinvoiceAmountCheckAPIResponse() *AlibabaEinvoiceAmountCheckAPIResponse {
	return poolAlibabaEinvoiceAmountCheckAPIResponse.Get().(*AlibabaEinvoiceAmountCheckAPIResponse)
}

// ReleaseAlibabaEinvoiceAmountCheckAPIResponse 将 AlibabaEinvoiceAmountCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceAmountCheckAPIResponse(v *AlibabaEinvoiceAmountCheckAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceAmountCheckAPIResponse.Put(v)
}

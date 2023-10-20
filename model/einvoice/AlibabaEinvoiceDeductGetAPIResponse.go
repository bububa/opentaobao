package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeductGetAPIResponse 发票扣减的接口 API返回值
// alibaba.einvoice.deduct.get
//
// 获取历史发票扣减量、每日发票扣减量的接口
type AlibabaEinvoiceDeductGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceDeductGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeductGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceDeductGetAPIResponseModel).Reset()
}

// AlibabaEinvoiceDeductGetAPIResponseModel is 发票扣减的接口 成功返回结果
type AlibabaEinvoiceDeductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_deduct_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaEinvoiceDeductGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeductGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceDeductGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceDeductGetAPIResponse)
	},
}

// GetAlibabaEinvoiceDeductGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceDeductGetAPIResponse
func GetAlibabaEinvoiceDeductGetAPIResponse() *AlibabaEinvoiceDeductGetAPIResponse {
	return poolAlibabaEinvoiceDeductGetAPIResponse.Get().(*AlibabaEinvoiceDeductGetAPIResponse)
}

// ReleaseAlibabaEinvoiceDeductGetAPIResponse 将 AlibabaEinvoiceDeductGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceDeductGetAPIResponse(v *AlibabaEinvoiceDeductGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceDeductGetAPIResponse.Put(v)
}

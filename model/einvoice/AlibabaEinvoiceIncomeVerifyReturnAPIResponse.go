package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeVerifyReturnAPIResponse 服务商回传发票查验的结果 API返回值
// alibaba.einvoice.income.verify.return
//
// 服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传
type AlibabaEinvoiceIncomeVerifyReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceIncomeVerifyReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeVerifyReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceIncomeVerifyReturnAPIResponseModel).Reset()
}

// AlibabaEinvoiceIncomeVerifyReturnAPIResponseModel is 服务商回传发票查验的结果 成功返回结果
type AlibabaEinvoiceIncomeVerifyReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_income_verify_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用结果，true=成功，false=失败，subCode以isp开头时需要服务商重试
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeVerifyReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceIncomeVerifyReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceIncomeVerifyReturnAPIResponse)
	},
}

// GetAlibabaEinvoiceIncomeVerifyReturnAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceIncomeVerifyReturnAPIResponse
func GetAlibabaEinvoiceIncomeVerifyReturnAPIResponse() *AlibabaEinvoiceIncomeVerifyReturnAPIResponse {
	return poolAlibabaEinvoiceIncomeVerifyReturnAPIResponse.Get().(*AlibabaEinvoiceIncomeVerifyReturnAPIResponse)
}

// ReleaseAlibabaEinvoiceIncomeVerifyReturnAPIResponse 将 AlibabaEinvoiceIncomeVerifyReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceIncomeVerifyReturnAPIResponse(v *AlibabaEinvoiceIncomeVerifyReturnAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceIncomeVerifyReturnAPIResponse.Put(v)
}

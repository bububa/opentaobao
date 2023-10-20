package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse 商家查询本次开票的保险单号 API返回值
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
type TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse struct {
	model.CommonResponse
	TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponseModel).Reset()
}

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponseModel is 商家查询本次开票的保险单号 成功返回结果
type TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponseModel struct {
	XMLName xml.Name `xml:"servindustry_finance_insurance_invoice_insurancenos_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse)
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse 从 sync.Pool 获取 TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse
func GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse() *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse {
	return poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse.Get().(*TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse 将 TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse 保存到 sync.Pool
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse(v *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse) {
	v.Reset()
	poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIResponse.Put(v)
}

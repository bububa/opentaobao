package txcs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTxcsFinanceInvoiceInputAPIResponse 供应商发票录入 API返回值
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
type TmallTxcsFinanceInvoiceInputAPIResponse struct {
	model.CommonResponse
	TmallTxcsFinanceInvoiceInputAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTxcsFinanceInvoiceInputAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTxcsFinanceInvoiceInputAPIResponseModel).Reset()
}

// TmallTxcsFinanceInvoiceInputAPIResponseModel is 供应商发票录入 成功返回结果
type TmallTxcsFinanceInvoiceInputAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_txcs_finance_invoice_input_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回内容
	Result *AccessBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTxcsFinanceInvoiceInputAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTxcsFinanceInvoiceInputAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTxcsFinanceInvoiceInputAPIResponse)
	},
}

// GetTmallTxcsFinanceInvoiceInputAPIResponse 从 sync.Pool 获取 TmallTxcsFinanceInvoiceInputAPIResponse
func GetTmallTxcsFinanceInvoiceInputAPIResponse() *TmallTxcsFinanceInvoiceInputAPIResponse {
	return poolTmallTxcsFinanceInvoiceInputAPIResponse.Get().(*TmallTxcsFinanceInvoiceInputAPIResponse)
}

// ReleaseTmallTxcsFinanceInvoiceInputAPIResponse 将 TmallTxcsFinanceInvoiceInputAPIResponse 保存到 sync.Pool
func ReleaseTmallTxcsFinanceInvoiceInputAPIResponse(v *TmallTxcsFinanceInvoiceInputAPIResponse) {
	v.Reset()
	poolTmallTxcsFinanceInvoiceInputAPIResponse.Put(v)
}

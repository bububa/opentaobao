package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperPrintAPIResponse 纸票打印接口 API返回值
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
type AlibabaEinvoicePaperPrintAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePaperPrintAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperPrintAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePaperPrintAPIResponseModel).Reset()
}

// AlibabaEinvoicePaperPrintAPIResponseModel is 纸票打印接口 成功返回结果
type AlibabaEinvoicePaperPrintAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果，打印结果tmc消息alibaba_invoice_PaperOpsReturn异步通知
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperPrintAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoicePaperPrintAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePaperPrintAPIResponse)
	},
}

// GetAlibabaEinvoicePaperPrintAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePaperPrintAPIResponse
func GetAlibabaEinvoicePaperPrintAPIResponse() *AlibabaEinvoicePaperPrintAPIResponse {
	return poolAlibabaEinvoicePaperPrintAPIResponse.Get().(*AlibabaEinvoicePaperPrintAPIResponse)
}

// ReleaseAlibabaEinvoicePaperPrintAPIResponse 将 AlibabaEinvoicePaperPrintAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePaperPrintAPIResponse(v *AlibabaEinvoicePaperPrintAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePaperPrintAPIResponse.Put(v)
}

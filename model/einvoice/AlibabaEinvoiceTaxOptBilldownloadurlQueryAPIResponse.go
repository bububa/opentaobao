package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse 税筹业务账单文件下载URL查询 API返回值
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
type AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponseModel is 税筹业务账单文件下载URL查询 成功返回结果
type AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_billdownloadurl_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse
func GetAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse() *AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse {
	return poolAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse.Get().(*AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse 将 AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse(v *AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse.Put(v)
}

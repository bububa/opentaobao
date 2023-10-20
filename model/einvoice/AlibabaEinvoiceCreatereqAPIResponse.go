package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCreatereqAPIResponse ERP开票请求接口 API返回值
// alibaba.einvoice.createreq
//
// ERP发起开票请求
type AlibabaEinvoiceCreatereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceCreatereqAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceCreatereqAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceCreatereqAPIResponseModel).Reset()
}

// AlibabaEinvoiceCreatereqAPIResponseModel is ERP开票请求接口 成功返回结果
type AlibabaEinvoiceCreatereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_createreq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票信息是否成功接受
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceCreatereqAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceCreatereqAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceCreatereqAPIResponse)
	},
}

// GetAlibabaEinvoiceCreatereqAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceCreatereqAPIResponse
func GetAlibabaEinvoiceCreatereqAPIResponse() *AlibabaEinvoiceCreatereqAPIResponse {
	return poolAlibabaEinvoiceCreatereqAPIResponse.Get().(*AlibabaEinvoiceCreatereqAPIResponse)
}

// ReleaseAlibabaEinvoiceCreatereqAPIResponse 将 AlibabaEinvoiceCreatereqAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceCreatereqAPIResponse(v *AlibabaEinvoiceCreatereqAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceCreatereqAPIResponse.Put(v)
}

package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceRedCreatereqAPIResponse 发票冲红接口 API返回值
// alibaba.einvoice.red.createreq
//
// 发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
type AlibabaEinvoiceRedCreatereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceRedCreatereqAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceRedCreatereqAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceRedCreatereqAPIResponseModel).Reset()
}

// AlibabaEinvoiceRedCreatereqAPIResponseModel is 发票冲红接口 成功返回结果
type AlibabaEinvoiceRedCreatereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_red_createreq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否冲红成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceRedCreatereqAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = ""
}

var poolAlibabaEinvoiceRedCreatereqAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceRedCreatereqAPIResponse)
	},
}

// GetAlibabaEinvoiceRedCreatereqAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceRedCreatereqAPIResponse
func GetAlibabaEinvoiceRedCreatereqAPIResponse() *AlibabaEinvoiceRedCreatereqAPIResponse {
	return poolAlibabaEinvoiceRedCreatereqAPIResponse.Get().(*AlibabaEinvoiceRedCreatereqAPIResponse)
}

// ReleaseAlibabaEinvoiceRedCreatereqAPIResponse 将 AlibabaEinvoiceRedCreatereqAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceRedCreatereqAPIResponse(v *AlibabaEinvoiceRedCreatereqAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceRedCreatereqAPIResponse.Put(v)
}

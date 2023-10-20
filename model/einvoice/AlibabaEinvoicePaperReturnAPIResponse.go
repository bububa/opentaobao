package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperReturnAPIResponse 纸质发票结果回传 API返回值
// alibaba.einvoice.paper.return
//
// 纸质发票结果回传
type AlibabaEinvoicePaperReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePaperReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePaperReturnAPIResponseModel).Reset()
}

// AlibabaEinvoicePaperReturnAPIResponseModel is 纸质发票结果回传 成功返回结果
type AlibabaEinvoicePaperReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务端接收开票回传数据的结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaEinvoicePaperReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePaperReturnAPIResponse)
	},
}

// GetAlibabaEinvoicePaperReturnAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePaperReturnAPIResponse
func GetAlibabaEinvoicePaperReturnAPIResponse() *AlibabaEinvoicePaperReturnAPIResponse {
	return poolAlibabaEinvoicePaperReturnAPIResponse.Get().(*AlibabaEinvoicePaperReturnAPIResponse)
}

// ReleaseAlibabaEinvoicePaperReturnAPIResponse 将 AlibabaEinvoicePaperReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePaperReturnAPIResponse(v *AlibabaEinvoicePaperReturnAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePaperReturnAPIResponse.Put(v)
}

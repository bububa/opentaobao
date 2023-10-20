package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperCommonReturnAPIResponse 纸票通用回传接口 API返回值
// alibaba.einvoice.paper.common.return
//
// 纸票通用回传接口（打印回传、注册回传等），只返回成功or失败
type AlibabaEinvoicePaperCommonReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePaperCommonReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperCommonReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePaperCommonReturnAPIResponseModel).Reset()
}

// AlibabaEinvoicePaperCommonReturnAPIResponseModel is 纸票通用回传接口 成功返回结果
type AlibabaEinvoicePaperCommonReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_common_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传接收结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperCommonReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoicePaperCommonReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePaperCommonReturnAPIResponse)
	},
}

// GetAlibabaEinvoicePaperCommonReturnAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePaperCommonReturnAPIResponse
func GetAlibabaEinvoicePaperCommonReturnAPIResponse() *AlibabaEinvoicePaperCommonReturnAPIResponse {
	return poolAlibabaEinvoicePaperCommonReturnAPIResponse.Get().(*AlibabaEinvoicePaperCommonReturnAPIResponse)
}

// ReleaseAlibabaEinvoicePaperCommonReturnAPIResponse 将 AlibabaEinvoicePaperCommonReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePaperCommonReturnAPIResponse(v *AlibabaEinvoicePaperCommonReturnAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePaperCommonReturnAPIResponse.Put(v)
}

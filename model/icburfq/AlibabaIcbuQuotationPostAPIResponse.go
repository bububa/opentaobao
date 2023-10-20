package icburfq

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuQuotationPostAPIResponse 供应商提交报价 API返回值
// alibaba.icbu.quotation.post
//
// 供应商对RFQ进行报价
type AlibabaIcbuQuotationPostAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuQuotationPostAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuQuotationPostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuQuotationPostAPIResponseModel).Reset()
}

// AlibabaIcbuQuotationPostAPIResponseModel is 供应商提交报价 成功返回结果
type AlibabaIcbuQuotationPostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_quotation_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回结果信息
	Result *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuQuotationPostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIcbuQuotationPostAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuQuotationPostAPIResponse)
	},
}

// GetAlibabaIcbuQuotationPostAPIResponse 从 sync.Pool 获取 AlibabaIcbuQuotationPostAPIResponse
func GetAlibabaIcbuQuotationPostAPIResponse() *AlibabaIcbuQuotationPostAPIResponse {
	return poolAlibabaIcbuQuotationPostAPIResponse.Get().(*AlibabaIcbuQuotationPostAPIResponse)
}

// ReleaseAlibabaIcbuQuotationPostAPIResponse 将 AlibabaIcbuQuotationPostAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuQuotationPostAPIResponse(v *AlibabaIcbuQuotationPostAPIResponse) {
	v.Reset()
	poolAlibabaIcbuQuotationPostAPIResponse.Put(v)
}
